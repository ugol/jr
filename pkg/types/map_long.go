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
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
	"io"
)

func writeMapLong(r map[string]int64, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteLong(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapLongWrapper struct {
	Target *map[string]int64
	keys   []string
	values []int64
}

func (_ *MapLongWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapLongWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapLongWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]int64, 0, s)
	}
}

func (r *MapLongWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapLongWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapLongWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v int64
	r.values = append(r.values, v)
	return &types.Long{Target: &r.values[len(r.values)-1]}
}

func (_ *MapLongWrapper) AppendArray() types.Field { panic("Unsupported operation") }
