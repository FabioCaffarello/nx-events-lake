import { InsertConfigsExecutorSchema } from './schema';
import executor from './executor';
import { ExecutorContext } from "@nrwl/devkit";

const options: InsertConfigsExecutorSchema = {
  source: 'test-source',
};

describe('InsertConfigs Executor', () => {
  it('can run', async () => {
    const output = await executor(options, {} as ExecutorContext);
    expect(output.success).toBe(true);
  });
});
