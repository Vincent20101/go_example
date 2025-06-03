package trace

import (
	"context"

	slog "gitlab.casa-systems.com/mobility/smf/infra/base/logger"
	pg "gitlab.casa-systems.com/mobility/smf/infra/base/pcap/grpc"

	"google.golang.org/grpc"
)

func SetPcapTraceOptAndContext(ctx context.Context, lc *slog.LogContext, s pg.ServiceName, rAddr pg.RemoteAddr, method string) (context.Context, []grpc.CallOption) {

	if !lc.IsNeedPcapTrace() {
		return ctx, []grpc.CallOption{}
	}

	forceCodec := pg.NewForceCodec(s, rAddr, lc.GetTracePcapContext().SubscriberId, lc, method)
	opt := grpc.ForceCodec(forceCodec)
	ctx = pg.ContextWithTraceInfo(ctx, forceCodec.GetCodecInfo())
	return ctx, []grpc.CallOption{opt}
}
