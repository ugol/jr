// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     clickstream.avsc
 *     clickstreamcodes.avsc
 *     clickstreamusers.avsc
 *     creditcards.avsc
 *     deviceinformation.avsc
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

type Pageviews struct {
	Viewtime int64 `json:"viewtime"`

	Userid string `json:"userid"`

	Pageid string `json:"pageid"`
}

const PageviewsAvroCRC64Fingerprint = "\xdbia\b=\x14\xae\xea"

func NewPageviews() Pageviews {
	r := Pageviews{}
	return r
}

func DeserializePageviews(r io.Reader) (Pageviews, error) {
	t := NewPageviews()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePageviewsFromSchema(r io.Reader, schema string) (Pageviews, error) {
	t := NewPageviews()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePageviews(r Pageviews, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Viewtime, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Userid, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Pageid, w)
	if err != nil {
		return err
	}
	return err
}

func (r Pageviews) Serialize(w io.Writer) error {
	return writePageviews(r, w)
}

func (r Pageviews) Schema() string {
	return "{\"fields\":[{\"name\":\"viewtime\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1,\"step\":10}},\"format_as_time\":\"unix_long\",\"type\":\"long\"}},{\"name\":\"userid\",\"type\":{\"arg.properties\":{\"regex\":\"User_[1-9]\"},\"type\":\"string\"}},{\"name\":\"pageid\",\"type\":{\"arg.properties\":{\"regex\":\"Page_[1-9][0-9]\"},\"type\":\"string\"}}],\"name\":\"ksql.pageviews\",\"type\":\"record\"}"
}

func (r Pageviews) SchemaName() string {
	return "ksql.pageviews"
}

func (_ Pageviews) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Pageviews) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Pageviews) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Pageviews) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Pageviews) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Pageviews) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Pageviews) SetString(v string)   { panic("Unsupported operation") }
func (_ Pageviews) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Pageviews) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Viewtime}

		return w

	case 1:
		w := types.String{Target: &r.Userid}

		return w

	case 2:
		w := types.String{Target: &r.Pageid}

		return w

	}
	panic("Unknown field index")
}

func (r *Pageviews) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Pageviews) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Pageviews) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Pageviews) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Pageviews) HintSize(int)                     { panic("Unsupported operation") }
func (_ Pageviews) Finalize()                        {}

func (_ Pageviews) AvroCRC64Fingerprint() []byte {
	return []byte(PageviewsAvroCRC64Fingerprint)
}

func (r Pageviews) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["viewtime"], err = json.Marshal(r.Viewtime)
	if err != nil {
		return nil, err
	}
	output["userid"], err = json.Marshal(r.Userid)
	if err != nil {
		return nil, err
	}
	output["pageid"], err = json.Marshal(r.Pageid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Pageviews) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["viewtime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Viewtime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for viewtime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["userid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Userid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for userid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["pageid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pageid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pageid")
	}
	return nil
}
