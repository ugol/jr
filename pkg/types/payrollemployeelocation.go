// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     clickstream.avsc
 *     clickstreamcodes.avsc
 *     clickstreamusers.avsc
 *     creditcards.avsc
 *     deviceinformation.avsc
 *     fleetmgmtdescription.avsc
 *     fleetmgmtlocation.avsc
 *     fleetmgmtsensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     netdevice.avsc
 *     orders.avsc
 *     pageviews.avsc
 *     payrollbonus.avsc
 *     payrollemployee.avsc
 *     payrollemployeelocation.avsc
 *     pizzaorders.avsc
 *     pizzaorderscancelled.avsc
 *     pizzaorderscompleted.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
 *     stockTrades.avsc
 *     stores.avsc
 *     sysloglogs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Payrollemployeelocation struct {
	Employee_id int32 `json:"employee_id"`

	Lab string `json:"lab"`

	Department_id int32 `json:"department_id"`

	Arrival_date int32 `json:"arrival_date"`
}

const PayrollemployeelocationAvroCRC64Fingerprint = "\x04tZ\xd4oƉ\x1a"

func NewPayrollemployeelocation() Payrollemployeelocation {
	r := Payrollemployeelocation{}
	return r
}

func DeserializePayrollemployeelocation(r io.Reader) (Payrollemployeelocation, error) {
	t := NewPayrollemployeelocation()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePayrollemployeelocationFromSchema(r io.Reader, schema string) (Payrollemployeelocation, error) {
	t := NewPayrollemployeelocation()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePayrollemployeelocation(r Payrollemployeelocation, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Employee_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Lab, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Department_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Arrival_date, w)
	if err != nil {
		return err
	}
	return err
}

func (r Payrollemployeelocation) Serialize(w io.Writer) error {
	return writePayrollemployeelocation(r, w)
}

func (r Payrollemployeelocation) Schema() string {
	return "{\"fields\":[{\"name\":\"employee_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1100,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"lab\",\"type\":{\"arg.properties\":{\"regex\":\"lab-\\\\d{1}\"},\"type\":\"string\"}},{\"name\":\"department_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":1}},\"type\":\"int\"}},{\"name\":\"arrival_date\",\"type\":{\"arg.properties\":{\"range\":{\"max\":19000,\"min\":18000}},\"logicalType\":\"date\",\"type\":\"int\"}}],\"name\":\"payroll.payrollemployeelocation\",\"type\":\"record\"}"
}

func (r Payrollemployeelocation) SchemaName() string {
	return "payroll.payrollemployeelocation"
}

func (_ Payrollemployeelocation) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetString(v string)   { panic("Unsupported operation") }
func (_ Payrollemployeelocation) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Payrollemployeelocation) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Employee_id}

		return w

	case 1:
		w := types.String{Target: &r.Lab}

		return w

	case 2:
		w := types.Int{Target: &r.Department_id}

		return w

	case 3:
		w := types.Int{Target: &r.Arrival_date}

		return w

	}
	panic("Unknown field index")
}

func (r *Payrollemployeelocation) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Payrollemployeelocation) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Payrollemployeelocation) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Payrollemployeelocation) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Payrollemployeelocation) HintSize(int)                     { panic("Unsupported operation") }
func (_ Payrollemployeelocation) Finalize()                        {}

func (_ Payrollemployeelocation) AvroCRC64Fingerprint() []byte {
	return []byte(PayrollemployeelocationAvroCRC64Fingerprint)
}

func (r Payrollemployeelocation) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["employee_id"], err = json.Marshal(r.Employee_id)
	if err != nil {
		return nil, err
	}
	output["lab"], err = json.Marshal(r.Lab)
	if err != nil {
		return nil, err
	}
	output["department_id"], err = json.Marshal(r.Department_id)
	if err != nil {
		return nil, err
	}
	output["arrival_date"], err = json.Marshal(r.Arrival_date)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Payrollemployeelocation) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["employee_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Employee_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for employee_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["lab"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lab); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for lab")
	}
	val = func() json.RawMessage {
		if v, ok := fields["department_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Department_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for department_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["arrival_date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Arrival_date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for arrival_date")
	}
	return nil
}
