# AWS Info

Gets information about a local machine on AWS by calling the AWS meta data
service.

The following URLs are called:

* *Instance info*: http://169.254.169.254/latest/dynamic/instance-identity/document

* *Public hostname*: http://169.254.169.254/latest/meta-data/public-hostname

## Usage

```
info, err := awsinfo.Get()
  if err != nil {
    return err
  }
fmt.Printf(info)
```

### Output

```
{
  "accountId" : "691722295028",
  "architecture" : "x86_64",
  "availabilityZone" : "eu-west-1a",
  "billingProducts" : null,
  "devpayProductCodes" : null,
  "imageId" : "ami-602cdc17",
  "instanceId" : "i-fee3b8bf",
  "instanceType" : "t1.micro",
  "kernelId" : "aki-52a34525",
  "pendingTime" : "2014-02-26T12:48:31Z",
  "privateIp" : "172.31.19.195",
  "publicHostname" : "ec2-52-18-51-59.eu-west-1.compute.amazonaws.com",
  "ramdiskId" : null,
  "region" : "eu-west-1",
  "version" : "2010-08-31"
}

