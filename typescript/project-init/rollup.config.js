// import commonjs from '@rollup/plugin-commonjs'
// import multiEntry from '@rollup/plugin-multi-entry';
// import babel from '@rollup/plugin-babel';
const input = ["build/src/index.js"];

export default [
    {
        input,
        plugins: [
            // commonjs(),
            // multiEntry(),
            // babel(),
        ],
        output: [
            {
                file: "./test/dist/rollup-bundle.js",
                format: "umd",
                name: "MyModule"
            }
        ]
    },
];