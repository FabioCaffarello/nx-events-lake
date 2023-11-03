import { CreateBucketExecutorSchema } from './schema';
import executor from './executor';

const options: CreateBucketExecutorSchema = {
  name: 'test-bucket',
};

describe('CreateBucket Executor', () => {
  it('can run', async () => {
    const output = await executor(options);
    expect(output.success).toBe(true);
  });
});
