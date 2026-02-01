# Ralph Wiggum OODA Loop

## TL;DR

Quick command reference showing the complete workflow from bootstrap to refactoring.

## What You Get

High-level feature overview of procedures, architecture, and quality control.

## How It Works

Core mechanics of the loop system.

### The Loop

Single iteration lifecycle and why fresh context matters.

### Planning vs Building

Two primary loop modes: planning procedures generate work, building procedures execute it.

### AGENTS.md: The Agent-Project Interface

Defines how agents interact with the project (work tracking, build commands, spec/impl definitions, quality criteria).

### OODA Phases: Breaking Down the Monolithic Prompt

The prompt to each loop iteration decomposes into four phases (observe, orient, decide, act).

### Composable Architecture

How prompt files combine into procedures.

#### Shared Components Across Procedures

Same OODA phase components reused in different procedure combinations.

#### The 8 Procedures

Brief overview of the purpose of each procedure (table format).

#### Custom Procedures

Creating your own procedure compositions.

## Key Principles

Critical patterns for agent behavior and quality control mechanisms.

### AGENTS.md as Source of Truth

Assumed inaccurate until verified empirically; updated when errors discovered.

### Skepticism and Empirical Verification

Don't assume not implemented; always search codebase first.

### Other Core Principles

Backpressure, parallel subagents, capture the why, tight tasks.

## Sample Repository Structure

Example file tree showing typical layout.

## Safety

Sandboxing requirements and blast radius philosophy.

## Troubleshooting

Common issues and escape hatches.

### Common Issues

Repeated implementations, test failures, missing code, off-track plans.

## Learn More

Links to detailed documentation (OODA Loop, Ralph Loop, Specs, etc.).
