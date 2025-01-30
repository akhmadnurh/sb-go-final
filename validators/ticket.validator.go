package validators

type FindTicketsQuery struct {
	PaginationQuery
}

type CreateTicketBody struct {
	EventID uint `json:"event_id" validate:"required"`
}

type CheckInBody struct {
	EventID    uint   `json:"event_id" validate:"required"`
	TicketCode string `json:"ticket_code" validate:"required,min=1,max=32"`
}
