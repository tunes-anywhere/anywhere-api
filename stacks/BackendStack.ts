import * as sst from "@serverless-stack/resources";
import { aws_certificatemanager, RemovalPolicy } from "aws-cdk-lib";
import { config } from "./config";

export const BackendStack = ({ stack, app }: sst.StackContext) => {
  const domainName = `anywhere-${stack.stage}.${config.awsDomainName}`;
  const autoDeleteObjects = app.local;
  const removalPolicy = app.local
    ? RemovalPolicy.DESTROY
    : RemovalPolicy.RETAIN;

  console.log(stack.stackName, {
    domainName,
    autoDeleteObjects,
    removalPolicy,
  });

  const bucket = new sst.Bucket(stack, "Bucket", {
    cdk: { bucket: { removalPolicy, autoDeleteObjects } },
    notifications: {
      // TODO notify ws/webhooks with changes
    },
  });

  stack.addDefaultFunctionEnv({ BUCKET_NAME: bucket.bucketName });
  stack.addDefaultFunctionPermissions([bucket]);

  const table = new sst.Table(stack, "Table", {
    fields: { pk: "string", sk: "string" },
    primaryIndex: { partitionKey: "pk", sortKey: "sk" },
    cdk: { table: { removalPolicy } },
    consumers: {
      // TODO download song and upload to bucket?
      // downloader: "handler/download_track/main.go",
      // TODO deletes orphaned songs
      // deleter: "handler/delete_track/main.go",
    },
  });

  // TODO real-time discovery/sync
  // const sync = new sst.WebSocketApi(stack, "Sync");

  // TODO dlq and error reporting
  // const dlq = new sst.Queue(stack, "DLQ", {
  //   // TODO failure notifications?
  //   // consumer: "handler/notify_failure/main.go"
  // });

  // const hostedZone = aws_route53.HostedZone.fromHostedZoneAttributes(
  //   stack,
  //   "HostedZone",
  //   { hostedZoneId: config.awsHostedZoneId, zoneName: config.awsSubdomainName }
  // );

  const certificate = aws_certificatemanager.Certificate.fromCertificateArn(
    stack,
    "Certificate",
    config.awsAcmCertificateArn
  );

  const api = new sst.Api(stack, "Api", {
    customDomain: { domainName, cdk: { certificate } },
    // TODO ApiKeyAuthorizer
    routes: {
      "GET /api/tracks": "handler/list_tracks/main.go",
      "GET /api/tracks/{id}": "handler/get_track/main.go",
      // TODO downloads song metadata for prefill?
      // "GET /api/metadata/{id}": "handler/download_track_metadata/main.go"
    },
  });

  stack.addOutputs({
    BucketName: bucket.bucketName,
    ApiUrl: api.url,
    DomainName: domainName,
  });

  return { bucket, table, api };
};
