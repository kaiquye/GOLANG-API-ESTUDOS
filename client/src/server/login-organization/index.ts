import { ServerAdapter } from "../adapter/server-adapter";
import {
  ILoginOrganizationIn,
  ILoginOrganizationOut,
} from "./interfaces/login-organization.interface";
import Api from "../api-instance";

class LoginOrganizationServer extends ServerAdapter<
  ILoginOrganizationIn,
  Promise<ILoginOrganizationOut>
> {
  async Run(data: ILoginOrganizationIn): Promise<ILoginOrganizationOut> {
    const response = await Api.post("/organization/login", {
      Document: data.document,
      Password: data.password,
    });

    return {
      error: this.parserResponse(response).error,
      accessToken: response.data?.Data.accessToken ?? "",
    };
  }
}
export default new LoginOrganizationServer();
