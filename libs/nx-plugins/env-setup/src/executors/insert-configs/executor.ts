import * as fs from "fs";
import * as path from "path";
import axios from "axios";
import { ExecutorContext } from "@nrwl/devkit";
import { InsertConfigsExecutorSchema } from "./schema";

const CONFIG_HANDLER_ENDPOINT = "http://localhost:8002/configs";

export default async function runExecutor(
  options: InsertConfigsExecutorSchema,
  context: ExecutorContext,
) {
  console.log("Executor ran for InsertConfigs", options);
  const configDir = path.join(context.root, ".configs", options.source);
  try {
    const configFiles = fs.readdirSync(configDir).filter((file) => file.endsWith("-config.json"));
    for (const configFile of configFiles) {
      const configPath = path.join(configDir, configFile);
      const jsonBody = JSON.parse(fs.readFileSync(configPath, "utf-8"));
      try {
        const response = await axios.post(CONFIG_HANDLER_ENDPOINT, jsonBody);
        console.log(`Successfully inserted config ${configFile}. Response: ${response.data}`);
      } catch (error) {
        console.log(`Error inserting config ${configFile}. Error: ${error}`);
      }
    }
  return {
    success: true,
  }
  } catch (error) {
    console.log(`Error reading configs from ${configDir}. Error: ${error}`);
    return {
      success: false,
    };
  }
}
