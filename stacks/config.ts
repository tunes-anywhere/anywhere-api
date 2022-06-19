const assertEnv = (key: string): string => {
  const value = process.env[key];

  if (!value) {
    throw new Error(`missing required env: "${key}"`);
  }

  return value;
};

export const config = Object.freeze({
  awsHostedZoneId: assertEnv("AWS_HOSTED_ZONE_ID"),
  awsDomainName: assertEnv("AWS_DOMAIN_NAME"),
  awsAcmCertificateArn: assertEnv("AWS_ACM_CERTIFICATE_ARN"),
  ytApiKey: assertEnv("YT_API_KEY"),
  debug: process.env.DEBUG === "true" ?? false,
});
