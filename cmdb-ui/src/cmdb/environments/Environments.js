import React from "react";
import {connect} from 'react-redux';

// from core components
import Table from "../../components/Table/Table";

// from Redux
import {listEnvironments, listLifecycles} from "../../actions";

// Table format is something like:
// <Table
//   tableHeaderColor="primary"
//   tableHead={["Name", "Country", "City", "Salary"]}
//   tableData={[
//     ["Dakota Rice", "Niger", "Oud-Turnhout", "$36,738"],
//     ["Minerva Hooper", "Curaçao", "Sinaai-Waas", "$23,789"],
//   ]}
// />

class Environments extends React.Component {
  componentDidMount() {
    this.props.listLifecycles();
    this.props.listEnvironments();
  }

   environmentTableData = () => {
     return this.props.environments.filter((environment) => {
      return (environment.name && environment.name.length);
    }).map((environment) => {
      const lifecycle = this.props.lifecycles[environment.lifecycle_id];
      if (lifecycle) {
        return [environment.name, lifecycle.name]
      }
      return [environment.name, ""];
    });
  };

  render() {
    return(
      <>
        <Table
          tableHeaderColor="primary"
          tableHead={["Name", "Lifecycle"]}
          tableData={this.environmentTableData()}
        />
      </>

    );
  }
}

const mapStateToProps = state => {
  return {
    // lifecycles: Object.values(state.lifecycles),
    lifecycles: state.lifecycles,
    environments: Object.values(state.environments),
  };
};

export default connect(mapStateToProps, { listLifecycles, listEnvironments })(Environments);
