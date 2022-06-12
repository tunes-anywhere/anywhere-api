const assertEnv = (key: string): string => {
  const value = process.env[key];

  if (!value) {
    throw new Error(`missing required env: "${key}"`);
  }

  return value;
};

const awsHostedZoneId = assertEnv("AWS_HOSTED_ZONE_ID");
const awsDomainName = assertEnv("AWS_DOMAIN_NAME");
const awsAcmCertificateArn = assertEnv("AWS_ACM_CERTIFICATE_ARN");

export const config = Object.freeze({
  awsHostedZoneId,
  awsDomainName,
  awsAcmCertificateArn,
});
