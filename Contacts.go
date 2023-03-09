package someRestApi

import "SomeRestApi/internal/handlers/hendler_gRPC/pb"

type Contacts struct {
	Id    int64  `json:"id"`
	Fio   string `json:"fio"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewContactFromPb(in *pb.Contact) *Contacts {
	return &Contacts{
		Id:    in.Id,
		Fio:   in.Fio,
		Email: in.Email,
		Phone: in.Phone,
	}
}

func NewContactToPb(out *Contacts) *pb.Contact {
	return &pb.Contact{
		Id:    out.Id,
		Fio:   out.Fio,
		Email: out.Email,
		Phone: out.Phone,
	}
}
