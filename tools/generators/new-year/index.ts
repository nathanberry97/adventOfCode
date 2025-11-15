import {
    Tree,
    formatFiles,
    generateFiles,
    joinPathFragments,
} from '@nx/devkit';

export default async function(tree: Tree, schema: any) {
    const year = String(schema.year);
    const projectRoot = `AoC/${year}`;

    tree.write(
        `${projectRoot}/project.json`,
        JSON.stringify(
            {
                name: `aoc-${year}`,
                sourceRoot: `${projectRoot}/src`,
                projectType: "application",
                targets: {
                    build: {
                        executor: "@nx/esbuild:esbuild",
                        outputs: [`{workspaceRoot}/build/AoC/${year}`],
                        options: {
                            outputPath: `build/AoC/${year}`,
                            format: ["cjs"],
                            bundle: true,
                            platform: "node",
                            main: `${projectRoot}/src/main.ts`,
                            tsConfig: `${projectRoot}/tsconfig.json`,
                            minify: false,
                            assets: [
                                {
                                    input: `${projectRoot}/src`,
                                    glob: "**/input.txt",
                                    output: ".",
                                },
                            ],
                        },
                    },
                    test: {
                        executor: "@nx/vite:test",
                        options: {
                            config: `${projectRoot}/vitest.config.ts`,
                            passWithNoTests: true,
                        },
                    },
                },
            },
            null,
            2
        )
    );

    generateFiles(tree, joinPathFragments(__dirname, './files'), projectRoot, { year });

    for (let i = 1; i <= 25; i++) {
        const day = i.toString().padStart(2, '0');
        const dayTemplatePath = joinPathFragments(__dirname, './templates/day');
        const dayTargetPath = joinPathFragments(projectRoot, `src/day/${day}`);
        generateFiles(tree, dayTemplatePath, dayTargetPath, { day });
    }

    await formatFiles(tree);
}
