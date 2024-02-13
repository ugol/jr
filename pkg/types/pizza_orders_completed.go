// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
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
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siem_logs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 *     users_array_map.avsc
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

type PizzaOrdersCompleted struct {
	Store_id int32 `json:"store_id"`

	Store_order_id int32 `json:"store_order_id"`

	Date int32 `json:"date"`

	Status string `json:"status"`

	Rack_time_secs int32 `json:"rack_time_secs"`

	Order_delivery_time_secs int32 `json:"order_delivery_time_secs"`
}

const PizzaOrdersCompletedAvroCRC64Fingerprint = "\xaa\xf7\xbf\fv\xe1\xadQ"

func NewPizzaOrdersCompleted() PizzaOrdersCompleted {
	r := PizzaOrdersCompleted{}
	return r
}

func DeserializePizzaOrdersCompleted(r io.Reader) (PizzaOrdersCompleted, error) {
	t := NewPizzaOrdersCompleted()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePizzaOrdersCompletedFromSchema(r io.Reader, schema string) (PizzaOrdersCompleted, error) {
	t := NewPizzaOrdersCompleted()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePizzaOrdersCompleted(r PizzaOrdersCompleted, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Store_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Store_order_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Date, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Status, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Rack_time_secs, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Order_delivery_time_secs, w)
	if err != nil {
		return err
	}
	return err
}

func (r PizzaOrdersCompleted) Serialize(w io.Writer) error {
	return writePizzaOrdersCompleted(r, w)
}

func (r PizzaOrdersCompleted) Schema() string {
	return "{\"fields\":[{\"name\":\"store_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":1}},\"type\":\"int\"}},{\"name\":\"store_order_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1000,\"step\":2}},\"type\":\"int\"}},{\"name\":\"date\",\"type\":{\"arg.properties\":{\"range\":{\"max\":19000,\"min\":18000}},\"logicalType\":\"date\",\"type\":\"int\"}},{\"name\":\"status\",\"type\":{\"arg.properties\":{\"options\":[\"completed\"]},\"type\":\"string\"}},{\"name\":\"rack_time_secs\",\"type\":{\"arg.properties\":{\"range\":{\"max\":230,\"min\":130}},\"type\":\"int\"}},{\"name\":\"order_delivery_time_secs\",\"type\":{\"arg.properties\":{\"range\":{\"max\":2000,\"min\":1100}},\"type\":\"int\"}}],\"name\":\"pizza_orders.PizzaOrdersCompleted\",\"type\":\"record\"}"
}

func (r PizzaOrdersCompleted) SchemaName() string {
	return "pizza_orders.PizzaOrdersCompleted"
}

func (_ PizzaOrdersCompleted) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetString(v string)   { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PizzaOrdersCompleted) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Store_id}

		return w

	case 1:
		w := types.Int{Target: &r.Store_order_id}

		return w

	case 2:
		w := types.Int{Target: &r.Date}

		return w

	case 3:
		w := types.String{Target: &r.Status}

		return w

	case 4:
		w := types.Int{Target: &r.Rack_time_secs}

		return w

	case 5:
		w := types.Int{Target: &r.Order_delivery_time_secs}

		return w

	}
	panic("Unknown field index")
}

func (r *PizzaOrdersCompleted) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PizzaOrdersCompleted) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PizzaOrdersCompleted) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) HintSize(int)                     { panic("Unsupported operation") }
func (_ PizzaOrdersCompleted) Finalize()                        {}

func (_ PizzaOrdersCompleted) AvroCRC64Fingerprint() []byte {
	return []byte(PizzaOrdersCompletedAvroCRC64Fingerprint)
}

func (r PizzaOrdersCompleted) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["store_id"], err = json.Marshal(r.Store_id)
	if err != nil {
		return nil, err
	}
	output["store_order_id"], err = json.Marshal(r.Store_order_id)
	if err != nil {
		return nil, err
	}
	output["date"], err = json.Marshal(r.Date)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	output["rack_time_secs"], err = json.Marshal(r.Rack_time_secs)
	if err != nil {
		return nil, err
	}
	output["order_delivery_time_secs"], err = json.Marshal(r.Order_delivery_time_secs)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PizzaOrdersCompleted) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["store_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["store_order_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_order_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_order_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for date")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	val = func() json.RawMessage {
		if v, ok := fields["rack_time_secs"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rack_time_secs); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rack_time_secs")
	}
	val = func() json.RawMessage {
		if v, ok := fields["order_delivery_time_secs"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Order_delivery_time_secs); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for order_delivery_time_secs")
	}
	return nil
}
