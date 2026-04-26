 # SITREP Generator

This application is a lightweight Go CLI that takes in a json file representing a conversation and returns a SITREP based on those messages.

The project was made to learn Go while also demonstrating a proof-of-concept for a SITREP generator.

---

## Features 
- Parse operational-style JSON message logs
- Filter low-value or irrelevant channel noise
- Generate concise SITREPs
- Extract key mission events
- Identify decisions and tasking
- Produce readable command-line summaries

---

## Examples
There are currently 3 example files:
- mq9sim.json simulates a recon mission and produces a SITREP.
- noisychat.json represents a conversation with some important elements but a lot of conversational "noise" that the API call needs to filter.
- example.json is a generic example of a conversation.

Examples are formatted as such:
```json
[
  { "time": "...", "channel": "...", "sender": "...", "message": "..." },
  { "time": "...", "channel": "...", "sender": "...", "message": "..." },
  { "time": "...", "channel": "...", "sender": "...", "message": "..." }
]
```

- The "time" field represents the time that the message was sent.
- The "channel" field represents the channel the message was sent to.
- The "sender" field represents who the sender of the message is.
- The "message" field is the actual text.

Once the user supplies a valid filename, a SITREP similar to this will be returned: 
```text
**SITREP – 26 April 2026**

**Timeline:**
- 1405Z: Blue Team reports Vehicle 12 maintenance delay; ETA return to service 45 minutes.
- 1408Z: Logistics confirms replacement vehicle ready if Vehicle 12 not operational by 1445Z.
- 1412Z: BattleCaptain delays convoy departure from 1430Z to 1515Z; orders update to section leads and movement board.
- 1420Z: All section leads acknowledge new departure time.
- 1432Z: Safety reports weather cell approaching route; advises monitoring visibility and reassessment at 1500Z.
- 1440Z: Intel confirms no external threat to convoy route.
- 1452Z: Vehicle 12 repair completed; asset available.
- 1455Z: BattleCaptain confirms maintaining 1515Z departure to preserve coordination.

**Summary:**
Vehicle 12 experienced a maintenance delay but was repaired ahead of the revised departure schedule. Convoy departure delayed 45 minutes, now set for 1515Z. Weather conditions potentially impacting route require monitoring with reassessment at 1500Z. No enemy threats reported.

**Suggestions for Action:**
- Conduct weather reassessment at 1500Z and decide on convoy movement accordingly.
- Maintain communications to ensure all units remain informed of any further changes.
- Prepare replacement vehicle on standby in case Vehicle 12 encounters issues before departure.
- Continue threat monitoring throughout convoy movement.
```

## Using the program
- Ensure that you have your OpenAI API key in an enviroment variable that is accessible by the program.
- Run the program via "go run ." when you are in the project root directory.
- When prompted to enter a file name, enter the name of one of the example files.


