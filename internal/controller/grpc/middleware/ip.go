package middleware

//func IP(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//	p, ok := peer.FromContext(ctx)
//	if !ok {
//		return nil, status.Error(codes.Code(401), "Failed to check ip address.")
//	}
//	ip := p.Addr.String()
//
//	if !isAllowed(ip) {
//		return nil, status.Error(codes.Code(401), "Client IP '"+ip+"' is not allowed.")
//	}
//	resp, err = handler(ctx, req)
//
//	return resp, err
//}
//
//func isAllowed(ip string) bool {
//	if strings.HasPrefix(ip, "[::1]") || strings.HasPrefix(ip, "127.0.0.1") {
//		return true
//	}
//	ip = strings.Split(ip, ":")[0]
//	octets := strings.Split(ip, ".")
//
//	whiteList := strings.Split(config.Get().AllowedIp, ";")
//
//Loop:
//	for _, allowedIP := range whiteList {
//		allowed := strings.Split(allowedIP, ".")
//		if len(allowed) != 4 {
//			continue Loop
//		}
//
//		for j := range octets {
//			if allowed[j] != "*" && allowed[j] != octets[j] {
//				continue Loop
//			}
//		}
//		return true
//	}
//
//	return false
//}
