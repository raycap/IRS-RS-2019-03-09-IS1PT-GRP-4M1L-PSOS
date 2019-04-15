import React, { Component } from 'react'
import { InputGroup, InputGroupAddon, InputGroupText, Input , Button , Form, FormGroup, Label , FormText , Table } from 'reactstrap'

class tableInput extends React.Component {
  render() {
    return (
		<Table bordered>
			<thead>
				<tr>
					<th></th>
					<th>Cost ($/unit)</th>
					<th>Price($/unit)</th>
					<th>Order Received (unit)</th>
				</tr>
			</thead>
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
			</tbody>
		</Table>
    )
  }
}

export default tableInput