import cmdb from "../api/cmdb";
import { LIST_APPLICATIONS } from "./types";

const headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU',
};

export const listApplications = () => async dispatch => {
  console.log('api call to get applications...')
  const response = await cmdb.get('/v1/applications', {headers});
  console.log(response.data)
  dispatch({type: LIST_APPLICATIONS, payload: response.data.results});
}

