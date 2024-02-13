// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
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
 *     stock_trades.avsc
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

type Payrollbonus struct {
	Employee_id int32 `json:"employee_id"`

	Bonus int32 `json:"bonus"`

	Ts int64 `json:"ts"`
}

const PayrollbonusAvroCRC64Fingerprint = "nN\x97\xf36\x05e$"

func NewPayrollbonus() Payrollbonus {
	r := Payrollbonus{}
	return r
}

func DeserializePayrollbonus(r io.Reader) (Payrollbonus, error) {
	t := NewPayrollbonus()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePayrollbonusFromSchema(r io.Reader, schema string) (Payrollbonus, error) {
	t := NewPayrollbonus()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePayrollbonus(r Payrollbonus, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Employee_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Bonus, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	return err
}

func (r Payrollbonus) Serialize(w io.Writer) error {
	return writePayrollbonus(r, w)
}

func (r Payrollbonus) Schema() string {
	return "{\"fields\":[{\"name\":\"employee_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1100,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"bonus\",\"type\":{\"arg.properties\":{\"range\":{\"max\":50,\"min\":10}},\"type\":\"int\"}},{\"name\":\"ts\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1640995199000,\"min\":1609459200000}},\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"payroll.payrollbonus\",\"type\":\"record\"}"
}

func (r Payrollbonus) SchemaName() string {
	return "payroll.payrollbonus"
}

func (_ Payrollbonus) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Payrollbonus) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Payrollbonus) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Payrollbonus) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Payrollbonus) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Payrollbonus) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Payrollbonus) SetString(v string)   { panic("Unsupported operation") }
func (_ Payrollbonus) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Payrollbonus) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Employee_id}

		return w

	case 1:
		w := types.Int{Target: &r.Bonus}

		return w

	case 2:
		w := types.Long{Target: &r.Ts}

		return w

	}
	panic("Unknown field index")
}

func (r *Payrollbonus) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Payrollbonus) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Payrollbonus) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Payrollbonus) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Payrollbonus) HintSize(int)                     { panic("Unsupported operation") }
func (_ Payrollbonus) Finalize()                        {}

func (_ Payrollbonus) AvroCRC64Fingerprint() []byte {
	return []byte(PayrollbonusAvroCRC64Fingerprint)
}

func (r Payrollbonus) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["employee_id"], err = json.Marshal(r.Employee_id)
	if err != nil {
		return nil, err
	}
	output["bonus"], err = json.Marshal(r.Bonus)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Payrollbonus) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["bonus"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bonus); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for bonus")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ts")
	}
	return nil
}
