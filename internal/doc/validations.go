package doc

import (
	"errors"

	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/l10n"
	"github.com/invopop/gobl/regimes/es"
)

var validSupplierLocalities = []l10n.Code{
	es.ZoneBI, // Vizcaya
	es.ZoneSS, // Guizpuzcoa
	es.ZoneVI, // Álava
}

func validate(inv *bill.Invoice, zone l10n.Code) error {
	if inv.Type == bill.InvoiceTypeCorrective {
		return errors.New("corrective invoices not supported, use credit or debit notes")
	}

	if inv.Supplier == nil || inv.Supplier.TaxID == nil {
		return nil // ignore
	}

	if zone == l10n.CodeEmpty {
		return errors.New("zone is required")
	}

	if !zone.In(validSupplierLocalities...) {
		return errors.New("zone not supported by TicketBAI")
	}

	if zone.In(es.ZoneSS, es.ZoneVI) {
		if len(inv.Lines) > 1000 {
			return errors.New("line count over limit (1000) for tax locality")
		}
		if inv.Customer != nil && len(inv.Customer.Addresses) == 0 {
			return errors.New("customer address required")
		}
	}

	return nil
}
