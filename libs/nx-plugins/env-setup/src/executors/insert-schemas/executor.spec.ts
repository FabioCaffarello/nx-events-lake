import { InsertSchemasExecutorSchema } from './schema';
import executor from './executor';
import { ExecutorContext } from "@nrwl/devkit";

const options: InsertSchemasExecutorSchema = {
  source: 'test-source',
};

describe('InsertSchemas Executor', () => {
  it('can run', async () => {
    const output = await executor(options as any, {} as ExecutorContext);
    expect(output.success).toBe(true);
  });
});
