package handler

import proto "srvs/userop_srv/proto/gen"

type UserOpServer struct {
	proto.UnimplementedAddressServer
	proto.UnimplementedUserFavServer
	proto.UnimplementedMessageServer
}
