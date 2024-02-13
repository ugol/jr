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

type Purchase struct {
	Id int64 `json:"id"`

	Item_type string `json:"item_type"`

	Quantity int64 `json:"quantity"`

	Price_per_unit Bytes `json:"price_per_unit"`
}

const PurchaseAvroCRC64Fingerprint = "Q\x81k\x93\xb7\xe9\xe4%"

func NewPurchase() Purchase {
	r := Purchase{}
	return r
}

func DeserializePurchase(r io.Reader) (Purchase, error) {
	t := NewPurchase()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePurchaseFromSchema(r io.Reader, schema string) (Purchase, error) {
	t := NewPurchase()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePurchase(r Purchase, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Item_type, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Quantity, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.Price_per_unit, w)
	if err != nil {
		return err
	}
	return err
}

func (r Purchase) Serialize(w io.Writer) error {
	return writePurchase(r, w)
}

func (r Purchase) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"long\"}},{\"name\":\"item_type\",\"type\":{\"arg.properties\":{\"options\":[\"boots\",\"shirt\",\"knife\",\"saucepan\",\"table\",\"chair\",\"clock\",\"frame\",\"guitar\",\"thermometer\",\"scarf\",\"doormat\",\"vase\",\"clippers\",\"game\",\"towels\"]},\"type\":\"string\"}},{\"name\":\"quantity\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":1}},\"type\":\"long\"}},{\"name\":\"price_per_unit\",\"type\":{\"arg.properties\":{\"range\":{\"max\":50,\"min\":10}},\"logicalType\":\"decimal\",\"precision\":4,\"scale\":2,\"type\":\"bytes\"}}],\"name\":\"datagen.example.purchase\",\"type\":\"record\"}"
}

func (r Purchase) SchemaName() string {
	return "datagen.example.purchase"
}

func (_ Purchase) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Purchase) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Purchase) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Purchase) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Purchase) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Purchase) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Purchase) SetString(v string)   { panic("Unsupported operation") }
func (_ Purchase) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Purchase) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Item_type}

		return w

	case 2:
		w := types.Long{Target: &r.Quantity}

		return w

	case 3:
		w := BytesWrapper{Target: &r.Price_per_unit}

		return w

	}
	panic("Unknown field index")
}

func (r *Purchase) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Purchase) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Purchase) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Purchase) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Purchase) HintSize(int)                     { panic("Unsupported operation") }
func (_ Purchase) Finalize()                        {}

func (_ Purchase) AvroCRC64Fingerprint() []byte {
	return []byte(PurchaseAvroCRC64Fingerprint)
}

func (r Purchase) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["item_type"], err = json.Marshal(r.Item_type)
	if err != nil {
		return nil, err
	}
	output["quantity"], err = json.Marshal(r.Quantity)
	if err != nil {
		return nil, err
	}
	output["price_per_unit"], err = json.Marshal(r.Price_per_unit)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Purchase) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["item_type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Item_type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for item_type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["quantity"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Quantity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for quantity")
	}
	val = func() json.RawMessage {
		if v, ok := fields["price_per_unit"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Price_per_unit); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for price_per_unit")
	}
	return nil
}
