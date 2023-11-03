import * as fs from "fs";
import * as path from "path";
import axios from "axios";
import { ExecutorContext } from "@nrwl/devkit";
import { InsertSchemasExecutorSchema } from './schema';

const SCHEMA_HANDLER_ENDPOINT = "http://localhost:8003/schemas";

export default async function runExecutor(
  options: InsertSchemasExecutorSchema,
  context: ExecutorContext,
) {
  console.log('Executor ran for InsertSchemas', options);
  const configDir = path.join(context.root, ".configs", options.source);
  try {
    const configFiles = fs.readdirSync(configDir).filter((file) => file.endsWith("-schema.json"));
    for (const configFile of configFiles) {
      const configPath = path.join(configDir, configFile);
      const jsonBody = JSON.parse(fs.readFileSync(configPath, "utf-8"));
      try {
        const response = await axios.post(SCHEMA_HANDLER_ENDPOINT, jsonBody);
        console.log(`Successfully inserted schema ${configFile}. Response: ${response.data}`);
      } catch (error) {
        console.log(`Error inserting schema ${configFile}. Error: ${error}`);
      }
    }
  return {
    success: true,
  }
  } catch (error) {
    console.log(`Error reading schemas from ${configDir}. Error: ${error}`);
    return {
      success: false,
    };
  }
}
