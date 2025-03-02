import React from "react";
import { connect } from 'react-redux';

// from core components
import EnvironmentSelects from "../environments/EnvironmentSelects";
import AppButton from "./AppButton";

// from Redux
import { listApplicationInstances, selectEnvironment } from "../../actions";
import Card from "../../components/Card/Card";
import CardHeader from "../../components/Card/CardHeader";
import CardIcon from "../../components/Card/CardIcon";
import IconButton from "@material-ui/core/IconButton";
import Cancel from "@material-ui/icons/Cancel";
import CardFooter from "../../components/Card/CardFooter";
import Update from "@material-ui/icons/Update";
import { makeStyles } from "@material-ui/core/styles";
import styles from "../../assets/jss/material-dashboard-react/views/dashboardStyle";
import {withStyles} from "@material-ui/core";

// Table format is something like:
// <Table
//   tableHeaderColor="primary"
//   tableHead={["Name", "Country", "City", "Salary"]}
//   tableData={[
//     ["Dakota Rice", "Niger", "Oud-Turnhout", "$36,738"],
//     ["Minerva Hooper", "Curaçao", "Sinaai-Waas", "$23,789"],
//   ]}
// />

const useStyles = makeStyles(styles);

class ApplicationInstances extends React.Component {

  constructor(props) {
    super(props);
    this.state = {showApp: false, appInstance: null};
  }


  selectEnvironment = (envId) => {
    this.props.selectEnvironment(envId);
    this.props.listApplicationInstances({envId});
  }

  // applicationInstanceTableData = () => {
  //   const appInstanceTableData = this.props.applicationInstances.filter( (applicationInstance) => {
  //     if (applicationInstance.name && applicationInstance.name.length ) {
  //       return true;
  //     } else {
  //       return false;
  //     }
  //   }).map( (applicationInstance) => {
  //     return [applicationInstance.name];
  //   });
  //
  //   return appInstanceTableData;
  // };
  //
  // render() {
  //   return(
  //     <>
  //       < EnvironmentSelects
  //         envName={this.state.envName}
  //         selectEnvironment={this.selectEnvironment}
  //       />
  //       <Table
  //         tableHeaderColor="primary"
  //         tableHead={["Name"]}
  //         tableData={this.applicationInstanceTableData()}
  //       />
  //     </>
  //
  //   );
  // }

  renderAppInstances = () => {
    return(
      <>
        <h3>{this.props.applicationInstances.length} Application Instances</h3>
        { this.props.applicationInstances.filter( (applicationInstance) => {
          return (applicationInstance.name && applicationInstance.name.length);
        }).map( (applicationInstance) => {
          return <AppButton key={applicationInstance.id} app={applicationInstance} onClick={() => {this.showAppView(applicationInstance)}}/>;
        })}
      </>
    );
  }

  renderConfigs = (configs) => {
    return configs.map( (config) => {
      return <div>{config.replace(/ /g, "\u00a0")}</div>;
     // return <div style={{whiteSpace: 'pre', color: 'black', background: 'pink'}}>{c}</div>;
    });
  }

  closeAppView = () => {
    this.setState({showApp: false})
  }

  showAppView = (applicationInstance) => {
    console.log("show app: ", applicationInstance);
    this.setState({showApp: true, applicationInstance: applicationInstance});
  }

  renderAppInstance = () => {
    console.log(this.state.applicationInstance.config_yaml);
    const configs = this.state.applicationInstance.config_yaml.split('\n');
    return(
      <div>
        <IconButton onClick={this.closeAppView}>
          <Cancel />
        </IconButton>
        <h3>{this.state.applicationInstance.name}</h3>
        <div>
          {this.renderConfigs(configs)}
        </div>
      </div>
    );
    // return(
    //   <Card>
    //     <CardHeader color="info" stats icon>
    //       <CardIcon color="info" onClick={this.closeAppView}>
    //         <Cancel />
    //       </CardIcon>
    //       <p className={this.props.classes.cardCategory}>{this.state.applicationInstance.config_yaml}</p>
    //       <h3 className={this.props.classes.cardTitle}>{this.state.applicationInstance.name}</h3>
    //     </CardHeader>
    //     <CardFooter stats>
    //       <div className={this.props.classes.stats}>
    //         <Update />
    //         Stats Here!
    //       </div>
    //     </CardFooter>
    //   </Card>
    //   );
  }

  render() {
    return(
      <>
        < EnvironmentSelects
          envId={this.props.envId}
          selectEnvironment={this.selectEnvironment}
        />
        {(this.state.showApp) ? this.renderAppInstance() : this.renderAppInstances()}
      </>
    );
  }
}

const mapStateToProps = state => {
  return {
    envId: state.selectedEnvId,
    applicationInstances: Object.values(state.applicationInstances),
  };
};

export default connect(mapStateToProps,
  {listApplicationInstances,
    selectEnvironment})
(withStyles(useStyles)(ApplicationInstances));
