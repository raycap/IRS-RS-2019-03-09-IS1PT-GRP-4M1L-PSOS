import React from 'react'
import TableForm from './Table'
import { Input, Button, ButtonGroup , Form, FormGroup, Label , Spinner } from 'reactstrap'
import './main.css'
// import axios from 'axios'
// import Chart from 'react-google-charts';
// import ReactChartkick, { LineChart, PieChart } from 'react-chartkick'
// import Chart from 'chart.js'

// ReactChartkick.addAdapter(Chart)

//export default App
class Submission extends React.Component {
  constructor() {
  	super();
  	this.state = {
      showResult: false,
      schedule: [],
      totalProfit: 0,
      componentsLeft:[],
      isLoading: true,
      errors: null
    }
    this.handleSubmit = this.handleSubmit.bind(this);
	this.showHome = this.showHome.bind(this);
	this.onRadioBtnClick = this.onRadioBtnClick.bind(this);
  }

  handleSubmit(event) {
    event.preventDefault();
    const data = new FormData(event.target);
	  const json = Object.assign(...Array.from(data, ([x,y]) => ({[x]:y})));
	  const inputData = {
      "quickScan": this.state.rSelected,
      "components": [{
  		  "name": "C1", 
  		  "cost": parseInt(json.A1),
  		  "price": parseInt(json.A2),
  		  "desiredUnit": parseInt(json.A3, 10)
      }, {
        "name": "C2",
	      "cost": parseInt(json.B1),
	      "price": parseInt(json.B2),
        "desiredUnit": parseInt(json.B3,10)
      }, {
        "name": "C3",
	      "cost": parseInt(json.C1),
	      "price": parseInt(json.C2),
        "desiredUnit": parseInt(json.C3,10)
      }, {
        "name": "C4",
	      "cost": parseInt(json.D1),
	      "price": parseInt(json.D2),
        "desiredUnit": parseInt(json.D3,10)
      }, {
        "name": "C5",
	      "cost": parseInt(json.E1),
	      "price": parseInt(json.E2),
        "desiredUnit": parseInt(json.E3,10)
      }, {
        "name": "C6",
	      "cost": parseInt(json.F1),
	      "price": parseInt(json.F2),
        "desiredUnit": parseInt(json.F3,10)
      }, {
        "name": "C7",
	      "cost": parseInt(json.G1),
	      "price": parseInt(json.G2),
        "desiredUnit": parseInt(json.G3,10)
      }, {
        "name": "C8",
	      "cost": parseInt(json.H1),
	      "price": parseInt(json.H2),
        "desiredUnit": parseInt(json.H3,10)
      }]
    };
	  alert(JSON.stringify(inputData));
    this.setState({ showResult: true })
	  
    // axios.post('https://localhost:8080/solve', { json },{})
    //   .then(response => {
    //     console.log(response);
    //     console.log(response.data);
    //     this.setState({
    //       schedule: response.data.schedule,
    //       maxProfits: response.data.maxProfits,
    //       numOfComponents: response.data.numOfComponents,
    //       isLoading: false,
    //       showResult: true
    //     });
    //   })
    //   // If we catch any errors connecting, let's update accordingly
    //   .catch(error => this.setState({ error, isLoading: false, showResult: false }));

    fetch('http://localhost:8080/solve', {
      method: 'post',
      body: JSON.stringify(inputData),
      mode: 'no-cors'

    }).then(function(response) {
      console.log(response)
      console.log(response.data)
      this.setState({
        schedule: response.data.schedule,
        totalProfit: response.data.totalProfit,
        componentsLeft: response.data.componentsLeft,
        isLoading: false,
        showResult: true
      });
    }).catch(error => this.setState({ error, isLoading: false, showResult: false }));

  }
  
  showHome() {
    this.setState({ 
      showResult: false 
    });
  }
  
  onRadioBtnClick(rSelected) {
    this.setState({ rSelected });
  }
  
  render() {
    const { componentsLeft, isLoading, showResult } = this.state;  
    if (showResult === true) {
      return (
        <div className='content-container'>  
          <h1>Machine Scheduling Optimizer</h1>
          <div className='form-container'>
            <Form>
              <FormGroup>
    	        <h2>Result</h2>
              <React.Fragment>
              <div>
                {!isLoading ? (
                  // <h2> Total profit is {totalProfit} </h2> 
                  // display schedule
                  componentsLeft.map(component => {
                    const { name, desiredUnit } = component;
                    return (
                      <h3>Number of {name} not able to be produced: {desiredUnit}</h3>
                    );
                  })
                ) : (
				  <div>
                    <Spinner type="grow" color="primary" />
                    <Spinner type="grow" color="secondary" />
                    <Spinner type="grow" color="success" />
                    <Spinner type="grow" color="danger" />
                    <Spinner type="grow" color="warning" />
                    <Spinner type="grow" color="info" />
                    <Spinner type="grow" color="light" />
                    <Spinner type="grow" color="dark" />
                  </div>
                )}
              </div>
              </React.Fragment>
				<Button color="primary" size="sm" onClick={this.showHome}>Back</Button>
              </FormGroup>  
            </Form>
          </div>
        </div>
      )	  
    };
	
    // const lineData = [
    //   {"name":"M1", "data": {"1": 3, "2017-01-02": 4}},
    //   {"name":"M2", "data": {"2017-01-01": 5, "2017-01-02": 3}},
    //   {"name":"M3", "data": {"2017-01-01": 5, "2017-01-02": 3}},
    //   {"name":"M4", "data": {"2017-01-01": 5, "2017-01-02": 3}},
    //   {"name":"M5", "data": {"2017-01-01": 5, "2017-01-02": 3}}
    // ];
	  return (
      <div className='content-container'>  
        <h1>Machine Scheduling Optimizer</h1>
        <div className='form-container'>
          <Form onSubmit={this.handleSubmit}>
            <FormGroup>
              <Label for="companyName">Company Name</Label>
              <Input type="textarea" name="company" id="companyName" placeholder="Enter Company Name" />
              <div className='checkbox-container'>
			    <h5>Quick Scan</h5>
                <ButtonGroup>
                  <Button color="primary" onClick={() => this.onRadioBtnClick(true)} active={this.state.rSelected === true}>On</Button>
                  <Button color="primary" onClick={() => this.onRadioBtnClick(false)} active={this.state.rSelected === false}>Off</Button>
                </ButtonGroup>
              </div>
            </FormGroup>
            <FormGroup>
    	      <TableForm />
            </FormGroup>
            <Button type="submit">Submit</Button>
          </Form>
        </div>
      </div>
    );
  }
}

export default Submission