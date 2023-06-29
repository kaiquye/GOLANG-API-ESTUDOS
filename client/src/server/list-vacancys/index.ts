import { ServerAdapter } from "../adapter/server-adapter";
import Api from "../api-instance";

class ListVacancysServer extends ServerAdapter<void, any> {
  async Run(): Promise<any> {
    const response = await Api.get("/vacancy");
    return response.data?.Data || [];
  }
}

export default new ListVacancysServer();
