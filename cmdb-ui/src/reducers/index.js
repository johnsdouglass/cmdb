import { combineReducers } from "redux";

import applicationReducer from "./applicationReducers";
import applicationInstanceReducer from "./applicationInstanceReducers";
import lifecycleReducer from "./lifecycleReducers";
import environmentReducer from "./environmentReducers";
import selectEnvReducer from "./selectEnvReducer";
import chartVersionReducer from "./chartVersionReducer";

export default combineReducers( {
  selectedEnvId: selectEnvReducer,
  applications: applicationReducer,
  applicationInstances: applicationInstanceReducer,
  lifecycles: lifecycleReducer,
  environments: environmentReducer,
  chartVersions: chartVersionReducer,
});
