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
		<th scope="row">Component A</th>
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
		<th scope="row">Component B</th>
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
		<th scope="row">Component C</th>
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
		<th scope="row">Component D</th>
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
		<th scope="row">Component E</th>
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
		<th scope="row">Component F</th>
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
		<th scope="row">Component G</th>
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
		<th scope="row">Component H</th>
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