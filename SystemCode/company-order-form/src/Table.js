import React, { Component } from 'react'
import { Table } from 'reactstrap'

const TableHeader = () => {
  return (
	<thead>
	  <tr>
		<th></th>
		<th>Cost ($/unit)</th>
		<th>Price ($/unit)</th>
		<th>Order Received (unit)</th>
	  </tr>
	</thead>
  )
}

const TableBody = () => {
  return (
	<tbody>
	  <tr>
		<th scope="row">Samsung S9 Silicone Case (Coloured)</th>
		<td>
		  <input type="text" name="A1" />
		</td>
		<td>
		  <input type="text" name="A2" />
		</td>
		<td>
		  <input type="text" name="A3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Huawei P30 Clear Case</th>
		<td>
		  <input type="text" name="B1" />
		</td>
		<td>
		  <input type="text" name="B2" />
		</td>
		<td>
		  <input type="text" name="B3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Samsung Galaxy Tab A.10.1 Case</th>
		<td>
		  <input type="text" name="C1" />
		</td>
		<td>
		  <input type="text" name="C2" />
		</td>
		<td>
		  <input type="text" name="C3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Microsoft Surface Pro 5 Protective Case</th>
		<td>
		  <input type="text" name="D1" />
		</td>
		<td>
		  <input type="text" name="D2" />
		</td>
		<td>
		  <input type="text" name="D3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Iphone X Normal Case</th>
		<td>
		  <input type="text" name="E1" />
		</td>
		<td>
		  <input type="text" name="E2" />
		</td>
		<td>
		  <input type="text" name="E3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Iphone XS SE Case (Gold Colour)</th>
		<td>
		  <input type="text" name="F1" />
		</td>
		<td>
		  <input type="text" name="F2" />
		</td>
		<td>
		  <input type="text" name="F3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Iphone XS Colour</th>
		<td>
		  <input type="text" name="G1" />
		</td>
		<td>
		  <input type="text" name="G2" />
		</td>
		<td>
		  <input type="text" name="G3" />
		</td>
	  </tr>
	  <tr>
		<th scope="row">Ipad Pro 12.9 inch Case</th>
		<td>
		  <input type="text" name="H1" />
		</td>
		<td>
		  <input type="text" name="H2" />
		</td>
		<td>
		  <input type="text" name="H3" />
		</td>
	  </tr>
    </tbody>
  )
}

class TableForm extends Component {
  render() {
    return (
      <Table bordered>
        <TableHeader />
        <TableBody />
      </Table>
    )
  }
}

export default TableForm