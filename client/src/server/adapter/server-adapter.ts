import { AxiosResponse } from "axios";

interface IParserError {
  error: boolean;
  message: string;
  infos?: object;
}

export abstract class ServerAdapter<Input, Output> {
  abstract Run(data: Input): Output;

  parserResponse(error: AxiosResponse): IParserError {
    console.log(error);

    if (error.status == 200 || error.status == 201) {
      return {
        error: false,
        message: error.data.message,
      };
    }
    return {
      error: true,
      message:
        error?.["response"].data.Message ?? error?.["response"].data.message,
    };
  }
}
