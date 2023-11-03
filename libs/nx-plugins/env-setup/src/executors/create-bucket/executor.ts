import { CreateBucketExecutorSchema } from './schema';
import * as minio from 'minio';

export default async function runExecutor(options: CreateBucketExecutorSchema) {
  console.log('Executor ran for CreateBucket', options);
  const minioClient = new minio.Client({
    endPoint: "localhost",
    port: 9000,
    useSSL: false,
    accessKey: "minio-root-user",
    secretKey: "minio-root-password",
  });

  try {
    await minioClient.makeBucket(options.name, "us-east-1");
    console.log(`Bucket ${options.name} created successfully in 'us-east-1'.`);
    return {
      success: true,
    };
  } catch (err) {
    console.log(`Error creating bucket ${options.name}. Err: ${err}`);
    return {
      success: false,
    };
  }
}
