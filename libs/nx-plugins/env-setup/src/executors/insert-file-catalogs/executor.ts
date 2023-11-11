import * as fs from "fs";
import * as path from "path";
import axios from "axios";
import { ExecutorContext } from "@nrwl/devkit";
import { InsertFileCatalogsExecutorSchema } from './schema';

const FILE_CATALOG_HANDLER_ENDPOINT = "http://localhost:8006/file-catalog";

export default async function runExecutor(
  options: InsertFileCatalogsExecutorSchema,
  context: ExecutorContext,
) {
  console.log('Executor ran for InsertFileCatalogs', options);
  const configDir = path.join(context.root, ".configs", options.source);
  try {
    const configFiles = fs.readdirSync(configDir).filter((file) => file.endsWith("-file-catalog.json"));
    for (const configFile of configFiles) {
      const configPath = path.join(configDir, configFile);
      const jsonBody = JSON.parse(fs.readFileSync(configPath, "utf-8"));
      try {
        const response = await axios.post(FILE_CATALOG_HANDLER_ENDPOINT, jsonBody);
        console.log(`Successfully inserted file catalog ${configFile}. Response: ${response.data}`);
      } catch (error) {
        console.log(`Error inserting file catalog ${configFile}. Error: ${error}`);
      }
    }
  return {
    success: true,
  }
  } catch (error) {
    console.log(`Error reading file catalogs from ${configDir}. Error: ${error}`);
    return {
      success: false,
    };
  }
}
