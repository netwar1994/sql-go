package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/netwar1994/sql-go/cmd/bank/app/dto"
	"github.com/netwar1994/sql-go/pkg/card"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cardSvc *card.Service
	mux *http.ServeMux
	ctx context.Context
	conn *pgxpool.Conn
}

type DbError struct {
	Err error
}

func NewDbError(err error) *DbError {
	return &DbError{Err: err}
}

func (e DbError) Error() string {
	return fmt.Sprintf("db error: %s", e.Err.Error())
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux, ctx context.Context, conn *pgxpool.Conn) *Server {
	return &Server{cardSvc: cardSvc, mux: mux, ctx: ctx, conn: conn}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/getTransactions", s.getTransactions)
	s.mux.HandleFunc("/getMostOftenBought", s.mostOftenBought)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	extractId := r.URL.Query().Get("id")
	if extractId == "" {
		response(w, errors.New("user id unspecified"))
		return
	}
	userId, err := strconv.ParseInt(extractId, 10, 64)
	if err != nil {
		response(w, errors.New("user id unspecified "))
	}

	rows, err := s.conn.Query(s.ctx,
		`SELECT id, issuer, number, balance, status, owner_id
			FROM cards
			WHERE owner_id = $1`, userId,
		)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Println(NewDbError(err))
			response(w, NewDbError(err))
			return
		}
	}
	defer rows.Close()

	var cards []*dto.CardDTO
	for rows.Next() {
		card := &dto.CardDTO{}
		err = rows.Scan(&card.Id, &card.Issuer, &card.Number, &card.Balance, &card.Status,  &card.OwnerId)
		if err != nil {
			log.Println(err)
			response(w, err)
			return
		}
		cards = append(cards, card)
	}

	response(w, cards)
}

func (s *Server) getTransactions(w http.ResponseWriter, r *http.Request) {
	extractId := r.URL.Query().Get("id")
	if extractId == "" {
		response(w, errors.New("user id unspecified"))
		return
	}
	userId, err := strconv.ParseInt(extractId, 10, 64)
	if err != nil {
		response(w, errors.New("user id unspecified "))
	}

	rows, err := s.conn.Query(s.ctx,
		`SELECT id, card_id, sum, mcc_id, description, status, created
			FROM transactions
			WHERE card_id = $1`, userId,
	)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Println(NewDbError(err))
			response(w, NewDbError(err))
			return
		}
	}
	defer rows.Close()

	var transactions []*dto.TransactionsDTO
	for rows.Next() {
		transaction := &dto.TransactionsDTO{}
		err = rows.Scan(&transaction.Id, &transaction.CardId, &transaction.Sum, &transaction.MccId,
			&transaction.Description, &transaction.Status, &transaction.Created)
		if err != nil {
			log.Println(err)
			response(w, err)
			return
		}
		transactions = append(transactions, transaction)
	}

	response(w, transactions)
}

func (s *Server) mostOftenBought(w http.ResponseWriter, r *http.Request) {
	extractId := r.URL.Query().Get("id")
	if extractId == "" {
		response(w, errors.New("user id unspecified"))
		return
	}
	userId, err := strconv.ParseInt(extractId, 10, 64)
	if err != nil {
		response(w, errors.New("user id unspecified "))
	}

	rows, err := s.conn.Query(s.ctx,
		`SELECT t.mcc_id, COUNT(t.mcc_id) AS cnt, m.description
			FROM transactions t
			JOIN mcc m ON t.mcc_id = m.id
			JOIN cards c ON t.card_id = c.id
			WHERE c.owner_id = $1
			GROUP BY t.mcc_id, m.description
			ORDER BY cnt DESC
			LIMIT 1`, userId,
	)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Println(userId)
			log.Println(NewDbError(err))
			response(w, NewDbError(err))
			return
		}
	}
	defer rows.Close()

	var transactions []*dto.MostOftenBoughtDTO
	for rows.Next() {
		transaction := &dto.MostOftenBoughtDTO{}
		err = rows.Scan(&transaction.MCCId, &transaction.Count, &transaction.Description)
		if err != nil {
			log.Println(err)
			response(w, err)
			return
		}
		transactions = append(transactions, transaction)
	}

	response(w, transactions)
}

func response(w http.ResponseWriter, dtos interface{}) {
	rBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(rBody)
	if err != nil {
		log.Println(err)
		return
	}
}
