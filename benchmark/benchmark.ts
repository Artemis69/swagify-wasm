import Go from "./wasm_exec.js";
import { swagify as swagifyTS } from "https://deno.land/x/swagify@v1.1.2/mod.ts";

declare global {
  var swagify: typeof swagifyTS;
}

const go = new Go();

const wasmCode = await Deno.readFile("../dist/main.wasm");
const wasmModule = new WebAssembly.Module(wasmCode);
const instance = await WebAssembly.instantiate(wasmModule, go.importObject);

go.run(instance);

import {
  bench,
  runBenchmarks,
} from "https://deno.land/std@0.113.0/testing/bench.ts";

bench({
  name: "WASM Swagify",
  runs: 1e5,
  func(b): void {
    b.start();
    swagify("Artemiy Shuckin", {});
    b.stop();
  },
});

bench({
  name: "TS Swagify",
  runs: 1e5,
  func(b): void {
    b.start();
    swagifyTS("Artemiy Shuckin", {});
    b.stop();
  },
});

runBenchmarks();
