import * as sst from "@serverless-stack/resources";
import { RemovalPolicy } from "aws-cdk-lib";
import { Certificate } from "aws-cdk-lib/aws-certificatemanager";
import { Distribution } from "aws-cdk-lib/aws-cloudfront";
import { S3Origin } from "aws-cdk-lib/aws-cloudfront-origins";
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

  const certificate = Certificate.fromCertificateArn(
    stack,
    "Certificate",
    config.awsAcmCertificateArn
  );

  const distribution = new Distribution(stack, "Distribution", {
    defaultBehavior: { origin: new S3Origin(bucket.cdk.bucket) },
  });

  const api = new sst.Api(stack, "Api", {
    defaults: { function: { permissions: [bucket, table] } },
    customDomain: { domainName, cdk: { certificate } },
    // TODO ApiKeyAuthorizer
    routes: {
      "GET /api/tracks": "handler/list_tracks/main.go",
      "POST /api/tracks": "handler/create_track/main.go",
      // "GET /api/tracks/{id}": "handler/get_track/main.go",
      // "PATCH /api/tracks/{id}": "handler/update_track/main.go",
      // TODO artists
      // TODO albums
    },
  });

  stack.addOutputs({
    BucketName: bucket.bucketName,
    ApiUrl: api.url,
    DistributionUrl: distribution.distributionDomainName,
    DomainName: domainName,
  });

  return { bucket, table, distribution, api };
};
