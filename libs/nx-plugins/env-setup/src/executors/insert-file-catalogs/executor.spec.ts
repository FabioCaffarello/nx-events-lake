import { InsertFileCatalogsExecutorSchema } from './schema';
import executor from './executor';
import { ExecutorContext } from "@nrwl/devkit";

const options: InsertFileCatalogsExecutorSchema = {
  source: 'ceaf',
};

describe('InsertConfigs Executor', () => {
  beforeEach(() => (options.source = "ceaf"));
  it('can`t run', async () => {
    const context: ExecutorContext = {
      root: '.',
      projectName: 'your-project-name',
      targetName: 'your-target-name',
      configurationName: 'your-configuration-name',
      cwd: '',
      isVerbose: false
    };
    options.source = 'wrong-ceaf';
    const output = await executor(options, context);
    console.log(output); // Add this line to log the output value to the console
    expect(output.success).toBe(false);
  });

  it('can run', async () => {
    const context: ExecutorContext = {
      root: '.',
      cwd: '',
      isVerbose: false
    };

    const output = await executor(options, context);
    console.log(output); // Add this line to log the output value to the console
    expect(output.success).toBe(true);
  });
});

