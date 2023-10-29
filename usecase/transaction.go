package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"reflect"
	"time"
)

type TransactionUsecase interface {
	Create(transaction dto.CreateTransactionRequest) (string, error)
	GetAll() (interface{}, error)
	GetByInvoice(invoice string) (model.Transaction, error)
	CheckSeatValidity(ticketID uint) (bool, error)
}

type transactionUsecase struct {
	transactionRepository		repository.TransactionRepository
	ticketUsecase						TicketUsecase
	showUsecase							ShowUsecase
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, ticketUsecase TicketUsecase, showUsecase ShowUsecase) *transactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepo,
		ticketUsecase: ticketUsecase,
		showUsecase: showUsecase,
	}
}

func validateCreateTransactionRequest(req dto.CreateTransactionRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func (s *transactionUsecase) Create(transaction dto.CreateTransactionRequest) (string, error) {
	e := validateCreateTransactionRequest(transaction)
	
	if e!= nil {
		return "", e
	}

	var tickets []model.Ticket
	for _, seat := range transaction.SeatID {
		ticket := dto.CreateTicketRequest{
			ShowID: transaction.ShowID,
			SeatID: seat,
		}
		ticketID, _ := s.ticketUsecase.GetTicket(ticket)
		seatAvailable, _ := s.ticketUsecase.CheckSoldOut(ticketID.ID)
		if (!seatAvailable) {
			tickets = append(tickets, ticketID)
		} else {
			return "", errors.New("There's unavailable seat(s) in your transaction")
		}
	}

	var price int
	show, _ := s.showUsecase.Find(int(transaction.ShowID))
	price = show.Price * len(tickets)

	var ticket_code string
	ticket_code = generateInvoiceNumber()

	transactionData := model.Transaction{
		AccountID: transaction.AccountID,
		Tickets: tickets,
		Date: time.Now(),
		TotalPrice: price,
		Status: "paid",
		TicketCode: ticket_code,
	}

	errs := s.transactionRepository.Create(transactionData)

	if errs != nil {
		return "", errs
	}

	return ticket_code, nil
}

func (s *transactionUsecase) GetAll() (interface{}, error) {
	transactions, err := s.transactionRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func generateInvoiceNumber() string {
	timestamp := time.Now().Unix()
	random_num := rand.Intn(5000)
	invoiceNumber := fmt.Sprintf("%d-%d", timestamp, random_num)
	return invoiceNumber
}

func (s *transactionUsecase) GetByInvoice(invoice string) (model.Transaction, error) {
	transaction, err := s.transactionRepository.GetByInvoice(invoice)

	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (s *transactionUsecase) CheckSeatValidity(ticketID uint) (bool, error) {
	filled, err := s.transactionRepository.CheckSeatValidity(ticketID)

	if err != nil {
		return !filled, err
	}

	return !filled, nil
}
