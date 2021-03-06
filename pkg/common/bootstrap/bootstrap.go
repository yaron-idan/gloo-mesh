package bootstrap

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/go-logr/zapr"
	"github.com/solo-io/gloo-mesh/pkg/common/schemes"
	"github.com/solo-io/gloo-mesh/pkg/common/utils/stats"
	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster/watch"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log"
	zaputil "sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// required import to enable kube client-go auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type StartParameters struct {
	Ctx             context.Context
	MasterManager   manager.Manager
	McClient        multicluster.Client
	Clusters        multicluster.Interface
	SnapshotHistory *stats.SnapshotHistory
	// Reference to Settings object this controller uses.
	SettingsRef v1.ObjectRef
	// enable additional logging
	VerboseMode bool
}

type StartReconciler func(
	parameters StartParameters,
) error

// bootstrap options for starting discovery
type Options struct {
	// MetricsBindPort is the TCP port that the controller should bind to
	// for serving prometheus metrics.
	// It can be set to 0 to disable the metrics serving.
	MetricsBindPort uint32

	// MasterNamespace if specified restricts the Master manager's cache to watch objects in the desired namespace.
	// Defaults to all namespaces.
	//
	// Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace.
	MasterNamespace string

	// enables verbose mode
	VerboseMode bool

	// ManagementContext if specified read the KubeConfig for the management cluster from this context. Only applies when running out of cluster.
	ManagementContext string

	// Reference to the Settings object that the controller should use.
	SettingsRef v1.ObjectRef
}

// convenience function for setting these options via spf13 flags
func (opts *Options) AddToFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&opts.MasterNamespace, "namespace", "n", metav1.NamespaceAll, "if specified restricts the master manager's cache to watch objects in the desired namespace.")
	flags.Uint32Var(&opts.MetricsBindPort, "metrics-port", defaults.MetricsPort, "port on which to serve Prometheus metrics. set to 0 to disable")
	flags.BoolVar(&opts.VerboseMode, "verbose", true, "enables verbose/debug logging")
	flags.StringVar(&opts.ManagementContext, "context", "", "If specified, use this context from the selected KubeConfig to connect to the local (management) cluster.")
	flags.StringVar(&opts.SettingsRef.Name, "settings-name", defaults.DefaultSettingsName, "The name of the Settings object this controller should use.")
	flags.StringVar(&opts.SettingsRef.Namespace, "settings-namespace", defaults.DefaultPodNamespace, "The namespace of the Settings object this controller should use.")
}

// the mesh-discovery controller is the Kubernetes Controller/Operator
// which processes k8s storage events to produce
// discovered resources.
func Start(ctx context.Context, rootLogger string, start StartReconciler, opts Options) error {
	setupLogging(opts.VerboseMode)

	ctx = contextutils.WithLogger(ctx, rootLogger)
	mgr, err := makeMasterManager(opts)
	if err != nil {
		return err
	}

	snapshotHistory := stats.NewSnapshotHistory()

	stats.MustStartServerBackground(snapshotHistory, opts.MetricsBindPort)

	clusterWatcher := watch.NewClusterWatcher(ctx, manager.Options{
		Namespace: "", // TODO (ilackarms): support configuring specific watch namespaces on remote clusters
		Scheme:    mgr.GetScheme(),
	})

	mcClient := multicluster.NewClient(clusterWatcher)

	params := StartParameters{
		Ctx:             ctx,
		MasterManager:   mgr,
		McClient:        mcClient,
		Clusters:        clusterWatcher,
		SnapshotHistory: snapshotHistory,
		VerboseMode:     opts.VerboseMode,
		SettingsRef:     opts.SettingsRef,
	}

	if err := start(params); err != nil {
		return err
	}

	if err := clusterWatcher.Run(mgr); err != nil {
		return err
	}

	contextutils.LoggerFrom(ctx).Infof("starting manager with options %+v", opts)
	return mgr.Start(ctx.Done())
}

// get the manager for the local cluster; we will use this as our "master" cluster
func makeMasterManager(opts Options) (manager.Manager, error) {
	cfg, err := config.GetConfigWithContext(opts.ManagementContext)
	if err != nil {
		return nil, err
	}

	mgr, err := manager.New(cfg, manager.Options{
		Namespace:          opts.MasterNamespace, // TODO (ilackarms): support configuring multiple watch namespaces on master cluster
		MetricsBindAddress: "0",                  // serve metrics using custom stats server
	})
	if err != nil {
		return nil, err
	}

	if err := schemes.AddToScheme(mgr.GetScheme()); err != nil {
		return nil, err
	}
	return mgr, nil
}

func setupLogging(verboseMode bool) {
	level := zapcore.InfoLevel
	if verboseMode {
		level = zapcore.DebugLevel
	}
	atomicLevel := zap.NewAtomicLevelAt(level)
	baseLogger := zaputil.NewRaw(
		zaputil.Level(&atomicLevel),
		// Only set debug mode if specified. This will use a non-json (human readable) encoder which makes it impossible
		// to use any json parsing tools for the log. Should only be enabled explicitly
		zaputil.UseDevMode(verboseMode),
	)

	// klog
	zap.ReplaceGlobals(baseLogger)
	// controller-runtime
	log.SetLogger(zapr.NewLogger(baseLogger))
	// go-utils
	contextutils.SetFallbackLogger(baseLogger.Sugar())
}
