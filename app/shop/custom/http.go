package custom

import (
	"github.com/coderi421/gframework/gmicro/server/restserver"
	"github.com/coderi421/goshop/app/shop/custom/config"
)

// NewUserHTTPServer 创建一个http service
func NewCustomHTTPServer(conf *config.Config) (*restserver.Server, error) {
	//trace.InitAgent(trace.Options{
	//	Name:     conf.Telemetry.Name,
	//	Endpoint: conf.Telemetry.Endpoint,
	//	Sampler:  conf.Telemetry.Sampler,
	//	Batcher:  conf.Telemetry.Batcher,
	//})

	cRestServer := restserver.NewServer(
		restserver.WithPort(conf.Server.HttpPort),
		restserver.WithEnableProfiling(true),
		restserver.WithMiddlewares(conf.Server.Middlewares),
		restserver.WithMetrics(true),
	)
	//_ = tracerProvider()

	//	配置好路由
	initRouter(cRestServer)
	return cRestServer, nil
}

//var tp *otelsdktrace.TracerProvider
//
//// 初始化Provider
//func tracerProvider() error {
//	url := "http://127.0.0.1:14268/api/traces"
//	jexp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
//	if err != nil {
//		panic(err)
//	}
//
//	tp = otelsdktrace.NewTracerProvider(
//		otelsdktrace.WithBatcher(jexp),
//		otelsdktrace.WithResource(
//			resource.NewWithAttributes(
//				semconv.SchemaURL,
//				semconv.ServiceNameKey.String("shop-user"),
//				attribute.String("environment", "dev"),
//				attribute.Int("ID", 1),
//			),
//		),
//	)
//	otel.SetTracerProvider(tp)
//	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
//	return nil
//}
