# ⭐ 3. MVP SCOPE (SMALLEST ADDICTIVE VERSION)

Below is what you must build first.  
No features beyond this are required for a fun, viral game.

---

## A) MVP Gameplay

### The Map

- **30–40 hex tiles**

Every tile has:

- **owner** (player or neutral)
- **units** (0–X)

### Player Actions (every 2–3 seconds)

- **Expand**: claim adjacent neutral tile
- **Attack**: fight another player's tile
- **Fortify**: add units to one of your tiles

### Combat

Simple deterministic rule:

```
Attacker units – Defender units
Winner controls tile, leftover units remain
```

- Units regenerate automatically at a slow rate per owned tile

> That's all — incredibly intuitive.

---

## B) MVP Win Condition

**After 10 minutes:**

- The player with the most tiles wins

**OR** if one player controls **60%+** they win early

---

## C) MVP Lobby & Match Flow

- Join a public match instantly
- Or "Host private match" with a share link
- Players spawn on random starting tiles
- Countdown 5 seconds → match begins

> **This is crucial:** click → instant game.

---

## D) MVP Social + Virality Layer

- Global chat inside match
- Ping a tile ("Help me!")

### Temporary Truce (non-binding)

- A simple UI that says "Truce?"
- If accepted, both see a grey line between them
- But they can still attack → betrayal remains

> This creates politics with zero complexity.

---

## E) MVP Progression

**Extremely light:**

- XP after each match
- Level 1–20

Each level gives:

- Cosmetic color for your tiles
- Or border shape

> **NO stats upgrades. Pure cosmetics.**

---

## F) MVP Tech Spec (Go WASM + Backend)

### Frontend (Ebiten WASM)

- Simple hex grid rendering
- Tap/click for actions
- UI: 2–3 buttons only

### Backend

- Authoritative game state
- 200ms tick for movement/combat
- WebSocket updates to clients
- Snapshot every 1–2s for reconnection
- Stateless match servers (easy scaling)

### Database

- Player accounts (optional for MVP)
- Cosmetic unlocks
- Match history (only last 20 stored)

---

# ⭐ 4. THE CORE EMOTIONAL DESIGN

This is what makes Claimers addicting, not just "fun":

## Territory = Identity

Players psychologically attach to land they own.  
Losing a tile hurts → players fight harder.

## Emergent Alliances

No formal team system → drama emerges naturally.

## Social Manipulation

Players lie, negotiate, and form pacts — this is the gasoline.

## Betrayal Moments

These are the most shareable and emotional moments.

## Short Matches → Infinite Loop

12 minutes is long enough for narrative, short enough for addiction.

## Skill + Politics

Not just mechanics — social strategy matters.  
It gives smart players an edge beyond APM.

---

# ⭐ THE FULL PACKAGE FOR CLAIMERS

Here is the structured plan:

## MVP

- 1 map
- 3 actions
- simple combat
- chat
- minimal cosmetics
- private match invite
- public quick match
- scoreboard
- matchmaking by random

**Time to build MVP:** 4–8 weeks (solo dev).

---

## Post-MVP (Month 2–6)

- persistent accounts
- more map layouts
- better cosmetics
- clans
- seasonal rewards

---

## Long-Term (Year 1+)

- mobile version (same codebase via WASM)
- map editor for communities
- tournaments
- streamer mode
- spectating
- analytics and heatmaps for tuning
