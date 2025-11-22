# CLAIMERS — Game Design Document (GDD)
Version 1.0 — Core MVP + Long-Term Vision
Written for: www.claimers.fun
## 1. OVERVIEW
Game Title: Claimers
Genre: Real-time territorial strategy (lightweight, social, competitive)
Platform: Web (Go WASM + Go backend)
Session Length: 5–12 minutes
Players: 4–20 per match
Target Audience: Casual players, strategy lovers, social gamers, groups of friends
High-Level Description

Claimers is a real-time territory control game where players expand, attack, and form temporary alliances on a shared hex map. The rules are extremely simple, but the social tension, diplomacy, and betrayals create depth and addiction.

The game focuses on:

Territory dominance

Emergent alliances

Short, intense matches

Viral friend-invites

Light cosmetic progression

## 2. CORE GAMEPLAY
### 2.1 Objective

Control the most territory at the end of the match OR dominate 60%+ of the map to win early.

### 2.2 Map

Hexagonal grid

~30–40 tiles for MVP

Tile properties:

Owner (none / player)

Unit count (integer ≥ 0)

Adjacent tiles list

### 2.3 Player Actions

Players may perform one action every 2–3 seconds:

1. Expand

Claim an adjacent neutral tile.

Costs a fixed amount of units (e.g., 3 units).

If insufficient units, action fails.

2. Attack

Attack an enemy tile from one of your tiles.

Combat is deterministic:

AttackerUnits - DefenderUnits


If attacker ≥ defender, tile changes owner → leftover units transfer.

3. Fortify

Add +1 unit to a controlled tile.

Units also regenerate passively per tile owned (e.g., +1 per 5 seconds).

### 2.4 Unit Generation

Base regeneration per owned tile: e.g., +0.2 units per second.

Maximum units per tile: soft cap (e.g., 10–20).

Goal: simple, predictable tempo that encourages expansion.

### 2.5 Match Flow

Player joins public or private match.

Automatic placement on random starting tile.

5-second countdown.

Match runs for 10 minutes.

End-game scoring:

Most tiles wins.

Early victory at 60%+ map control.

## 3. SOCIAL SYSTEMS
### 3.1 In-Game Chat

Basic text chat visible to all players in the match.

Used for diplomacy, threats, alliances, and betrayals.

### 3.2 Temporary “Truce” System

Players may request a truce with another player.

If accepted:

A grey line connects their capitals.

No gameplay restrictions (can still attack).

Purely social / psychological feature.

This creates emergent politics without complexity.

### 3.3 Tile Ping (“Help Me!”)

Players can ping a tile they own:

Visible to all players.

Encourages alliances and coordinated attacks.

## 4. PROGRESSION & REWARDS
### 4.1 XP System (MVP)

Earn XP after each match based on:

Tiles owned

Match placement

Time survived

XP levels unlock purely cosmetic rewards.

### 4.2 Cosmetic Unlocks

Tile color variations

Border shapes (thin, thick, glowing, jagged)

Capital icons (crown, skull, flag, temple)

Not allowed in MVP:

No stat boosts

No gameplay advantages

## 5. USER EXPERIENCE (UX) FLOW
### 5.1 Main Menu

Play Now (quick join → public match)

Private Match

Generate invite link

Start game when ready

Cosmetics

Settings (sound, animations on/off)

Login (optional MVP)

### 5.2 Match UI

Map View: Hex grid

Action Buttons:

Expand

Attack

Fortify

Player Panel:

Name

Tile ownership count

XP gain

Chat Box

Game Timer

Tile Inspector (on hover/click)

### 5.3 Post-Match

Total tiles controlled

Match placement

XP gained

Cosmetic unlocks (if any)

## 6. GAME RULES AND BALANCE
6.1 Starting Conditions

Each player begins with:

1 tile

3 units

20s “protected” period where no one can attack them (optional MVP)

6.2 Tile Limits

To avoid “hoarding” and turtling:

Soft unit cap (e.g., 15) per tile.

Excess units force players to expand.

6.3 Combat Rules

100% deterministic, no randomness.

Encourages planning before attacking.

Easy for new players to understand instantly.

6.4 Player Elimination

If a player loses their last tile, they are removed.

Optional future version: become a spectator or vassal (not MVP).

## 7. TECHNICAL DESIGN (HIGH LEVEL)
### 7.1 Client (Go WASM + Ebiten)

Renders hex map at 60 fps

UI buttons + chat

Minimal animation (color change, pulse on attack)

### 7.2 Network

Client → Server: action requests

Server → Clients: authoritative state broadcasts

WebSocket sync

~200ms tick rate

Delta updates for performance

### 7.3 Backend (Go)
Responsibilities:

Map state management

Player actions validation

Combat resolution

Timer + match lifecycle

XP and cosmetic progression storage

Match logs (light)

Backend must remain authoritative to prevent cheating.

### 7.4 Server Architecture

Stateless match servers

One central orchestrator for:

Matchmaking

Player persistence

Cosmetics

Allows horizontal scaling.

## 8. ART STYLE & SOUND
Visual Style

Minimalistic

Clean color-coded tiles

Smooth borders

Vibrant, readable palette

Icons

Simple geometric symbols

Crown as default capital icon

Sound

Soft click feedback

Expand / attack / fortify subtle audio cues

Match start and end chime

(Keep minimal for first version.)

## 9. LONG-TERM ROADMAP
Phase 1 (1–2 months after MVP)

Accounts + persistent cosmetics

More cosmetic variety

Additional maps

Player statistics and match history

Phase 2 (3–6 months)

Clans

Seasonal ranks

Seasonal map themes

Vassal system (defeated players join dominator)

Phase 3 (6–12 months)

Leaderboards

Tournaments

Clan wars

Spectator mode

Mobile optimization

Streamer/creator tools

## 10. FUTURE EXPANSION IDEAS (OPTIONAL)

Fog of war

Special tile types (volcano, gold mine)

Draft mode

Map maker for community

Unique “capital abilities” per cosmetic set (purely visual, not gameplay)

## 11. MVP SCOPE SUMMARY
Included:

40-tile hex map

Expand / Attack / Fortify actions

Deterministic combat

Unit regen

Truce requests

Tile ping

Chat

XP + cosmetics

Public & private matches

WebSocket real-time sync

Simple UI

Not included (for now):

Clans

Mobile controls

Ranked play

Vassal system

Match persistence

Special events or tiles

Power-ups

Paid cosmetics

Keep it lean.

## 12. SUCCESS METRICS
Early KPIs to track:

Average session length

Repeat players

Invitations sent per player

Average match completion rate

Tiles per minute per player

% of matches with 3+ social interactions (chat / truce / ping)

These metrics validate addictiveness and virality.