import * as sst from "@serverless-stack/resources";

export class BackendStack extends sst.Stack {
  bucket = new sst.Bucket(this, "Bucket");

  table = new sst.Table(this, "Table", {
    fields: { pk: "string", sk: "string" },
    primaryIndex: { partitionKey: "pk", sortKey: "sk" },
  });

  api = new sst.Api(this, "Api");

  // TODO real-time discovery/sync
  // sync = new sst.WebSocketApi(thisd, "Sync");

  dlq = new sst.Queue(this, "DLQ");

  constructor(scope: sst.App, id: string, props?: sst.StackProps) {
    super(scope, id, props);

    scope.addDefaultFunctionEnv({
      BUCKET_NAME: this.bucket.bucketName,
    });

    scope.addDefaultFunctionPermissions([this.bucket]);

    this.api.addRoutes(this, {
      "GET /api/tracks": "handler/list_tracks/main.go",
      "GET /api/tracks/{id}": "handler/get_track/main.go",
      // TODO downloads song metadata for prefill?
      // "GET /api/metadata/{id}": "handler/download_track_metadata/main.go"
    });

    this.table.addConsumers(this, {
      // download song and upload to bucket?
      // downloader: "handler/download_track/main.go"
      // deletes orphaned songs
      // deleter: "handler/delete_track/main.go"
    });

    // TODO failure notifications?
    // this.dlq.addConsumer(this, "handler/notify_failure/main.go")

    this.addOutputs({
      BucketName: this.bucket.bucketName,
      ApiUrl: this.api.url,
    });
  }
}
