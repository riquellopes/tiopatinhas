package organizze

type TransactionsResponse []struct {
	ID                      int         `json:"id"`
	Description             string      `json:"description"`
	Date                    string      `json:"date"`
	Paid                    bool        `json:"paid"`
	AmountCents             int         `json:"amount_cents"`
	TotalInstallments       int         `json:"total_installments"`
	Installment             int         `json:"installment"`
	Recurring               bool        `json:"recurring"`
	AccountID               int         `json:"account_id"`
	AccountType             string      `json:"account_type"`
	CategoryID              int         `json:"category_id"`
	ContactID               interface{} `json:"contact_id"`
	Notes                   string      `json:"notes"`
	AttachmentsCount        int         `json:"attachments_count"`
	CreditCardID            int         `json:"credit_card_id"`
	CreditCardInvoiceID     int         `json:"credit_card_invoice_id"`
	PaidCreditCardID        interface{} `json:"paid_credit_card_id"`
	PaidCreditCardInvoiceID interface{} `json:"paid_credit_card_invoice_id"`
	OpositeTransactionID    interface{} `json:"oposite_transaction_id"`
	OpositeAccountID        interface{} `json:"oposite_account_id"`
	CreatedAt               string      `json:"created_at"`
	UpdatedAt               string      `json:"updated_at"`
}
