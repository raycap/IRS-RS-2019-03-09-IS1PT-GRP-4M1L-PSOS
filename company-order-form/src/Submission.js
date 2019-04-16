import React from 'react'
import TableForm from './Table'
import { Input, Button, ButtonGroup , Form, FormGroup, Label , Spinner } from 'reactstrap'
import './main.css'
import axios from 'axios'
import Chart from 'react-google-charts';

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
  		  "materialCost": parseInt(json.A1),
  		  "price": parseInt(json.A2),
  		  "desiredUnit": parseInt(json.A3, 10)
      }, {
        "name": "C2",
	      "materialCost": parseInt(json.B1),
	      "price": parseInt(json.B2),
        "desiredUnit": parseInt(json.B3,10)
      }, {
        "name": "C3",
	      "materialCost": parseInt(json.C1),
	      "price": parseInt(json.C2),
        "desiredUnit": parseInt(json.C3,10)
      }, {
        "name": "C4",
	      "materialCost": parseInt(json.D1),
	      "price": parseInt(json.D2),
        "desiredUnit": parseInt(json.D3,10)
      }, {
        "name": "C5",
	      "materialCost": parseInt(json.E1),
	      "price": parseInt(json.E2),
        "desiredUnit": parseInt(json.E3,10)
      }, {
        "name": "C6",
	      "materialCost": parseInt(json.F1),
	      "price": parseInt(json.F2),
        "desiredUnit": parseInt(json.F3,10)
      }, {
        "name": "C7",
	      "materialCost": parseInt(json.G1),
	      "price": parseInt(json.G2),
        "desiredUnit": parseInt(json.G3,10)
      }, {
        "name": "C8",
	      "materialCost": parseInt(json.H1),
	      "price": parseInt(json.H2),
        "desiredUnit": parseInt(json.H3,10)
      }]
    };
	  alert(JSON.stringify(inputData));
    this.setState({ showResult: true })

// mocked api server : https://8086ab03-0c27-451b-bfad-0f4424821753.mock.pstmn.io/solve
//     axios.post('https://8086ab03-0c27-451b-bfad-0f4424821753.mock.pstmn.io/solve', json, {
      axios.post('http://localhost:8080/solve',  inputData, {
      headers: {
        // 'Content-Type': 'application/json',
        // 'Access-Control-Allow-Origin': '*'
      }
    })
//     axios.post('http://localhost:8080/solve', { json })
      .then(response => {
        console.log(response);
        console.log(response.data);
        this.setState({
          batches: response.data.batches,
          totalProfit: response.data.totalProfit,
          componentsLeft: response.data.componentsLeft,
          isLoading: false,
          showResult: true
        });
      })
      // If we catch any errors connecting, let's update accordingly
      .catch(error => this.setState({ error, isLoading: false, showResult: false }));

    // fetch('http://localhost:8080/solve', {
    // // fetch('https://6a3fc180-0789-441a-ae49-679b017d51c7.mock.pstmn.io/solve', {
    //   method: 'post',
    //   body: JSON.stringify(inputData),
    //   mode: 'no-cors'
    //
    // }).then(function(response) {
    //   console.log(response)
    //   console.log(response.data)
    //   this.setState({
    //     batches: response.data.batches,
    //     totalProfit: response.data.totalProfit,
    //     componentsLeft: response.data.componentsLeft,
    //     isLoading: false,
    //     showResult: true
    //   });
    // }).catch(error => this.setState({ error, isLoading: false, showResult: false }));

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
    const { componentsLeft, isLoading, showResult, totalProfit, batches } = this.state;  

  
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
                  <div>
                    <p> Total profit is ${totalProfit}. </p> 
                    <ComponentsLeftList componentsLeft={componentsLeft} />
                    <BatchesList batches={batches}/>
                  </div>  
                    // componentsLeft.map(component => {
                    //   const { name, desiredUnit } = component;
                    //   return (

                    //     <h3>Number of {name} not able to be produced: {desiredUnit}</h3>
                    //   );
                    // })
                  // display schedule
                  // <TableForm />
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

function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}

function ListItem(props) {
    return <li>{props.name}: {props.value}</li>;
}

function ComponentsLeftList(props) {
  const componentsLeft = props.componentsLeft;
  const listItems = componentsLeft.map((component) =>
    <ListItem key={component.name} value={component.desiredUnit} name={component.name} />
  );
  return (
    <div>
      <h5>Components that cannot be produced within the month</h5>
      <ul>
        {listItems}
      </ul> 
    </div> 
  );
}

function BatchesList(props) {
  const batches = props.batches;
  const batchItems = batches.map((batch, index) =>
    <TimelinePerBatch key={index} number={index+1} batch={batch}/>
  );
  return (
    <div className='batch-container'>
      {batchItems}
    </div>
  );
}

function TimelinePerBatch(props) {
  const batch = props.batch
  const headerData = [
    { type: 'string', id: 'MachineName' },
    { type: 'string', id: 'ComponentName' },
    { type: 'string', role: 'tooltip' },
    { type: 'date', id: 'Start' },
    { type: 'date', id: 'End' }
  ]
  const data = GetBatchSchedule(batch.machineSchedules)
  data.unshift(headerData)
  return (
    <div>
      <h3> Batch No. {props.number} </h3>
      <Chart
      width={'100%'}
      height={'40%'}
      chartType="Timeline"
      loader={<div>Loading Chart</div>}
      data={data}
      options={{
        showRowNumber: true,
      }}
      rootProps={{ 'data-testid': '1' }}
      />
    </div>
  );  
}

function GetBatchSchedule(machineSchedules) {
  const machineArr = ['M1','M2','M3','M4','M5','M6'];
  let finalArr = [];


  machineArr.forEach((machine) => {
    let a = machineSchedules[machine];
    if (a) {
      a.forEach((details) => {
        let b = [machine, details['componentName'], details['processName'], add_minutes(new Date(0, 0, 0, 0, 0, 0), details['startTime']), add_minutes(new Date(0, 0, 0, 0, 0, 0), details['endTime'])]
        finalArr.push(b)
      });
    }
  });
  return finalArr
}

const add_minutes =  function (dt, minutes) {
    dt.setMinutes(minutes)
    return dt
}

export default Submission