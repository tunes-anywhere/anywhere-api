import * as sst from "@serverless-stack/resources";
import { BackendStack } from "./BackendStack";

export default function main(app: sst.App) {
  app.setDefaultFunctionProps({
    runtime: "go1.x",
    environment: {
      CGO_ENABLED: "0",
      GOOS: "linux",
      GOARCH: "amd64",
    },
  });

  new BackendStack(app, "BackendStack");
}
