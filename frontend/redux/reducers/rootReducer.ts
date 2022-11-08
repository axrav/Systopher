import { combineReducers } from "redux";
import authReducer from "./authReducer";

const rootReducer = combineReducers({
  main: () => "main",
  auth: authReducer,
});

export default rootReducer;
