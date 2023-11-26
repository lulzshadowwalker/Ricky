package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var frameCount int; 
var mab *Map;

type Player struct {
  movingSprite rl.Texture2D;
  rollingSprite rl.Texture2D;
  slashingSprite rl.Texture2D;
  spriteFrame int; 
  src rl.Rectangle;
  dest rl.Rectangle; 
  pos rl.Vector2;
  dir MovementDirection;
  speed float32; 
  movementState MovementState;
}

type MovementState int
const (
  movementStateIdle MovementState = iota;
  movementStateWalking;
  movementStateRolling; 
  movementStateSlashing; 
)

func NewPlayer(movingSprite rl.Texture2D, rollingSprite rl.Texture2D, slashingSprite rl.Texture2D, src rl.Rectangle, dest rl.Rectangle, pos rl.Vector2, speed float32) *Player {
  return &Player{
    movingSprite: movingSprite, 
    rollingSprite: rollingSprite,
    slashingSprite: slashingSprite,
    src: src, 
    dest: dest, 
    pos: pos,
    speed: speed,
    movementState: movementStateWalking,
  }
}


func (p *Player) isMidAnimation() bool {
  switch player.movementState {
  case movementStateRolling, movementStateSlashing: return true;
  default: return false; 
  }
}

func (p *Player) setDirection(dir MovementDirection) {
  if p.isMidAnimation() { return; }

  if dir == movingIdle {
    p.dir = 0;
    return;
  }

  p.dir = p.dir | dir;
}

func (p *Player) roll() {
  player.movementState = movementStateRolling;
  // TODO, vector normalization for diagonal movement

  var ySpriteOffset float32;
  var vel rl.Vector2;
  speed := 0.85 * player.speed;

  switch player.dir {
  case movingIdle, movingDown: 
    ySpriteOffset = 0;
    vel = rl.NewVector2(0, -speed);
  case movingDownLeft: 
    ySpriteOffset = 1;
    vel = rl.NewVector2(speed, -speed);
  case movingLeft:
    ySpriteOffset = 2
    vel = rl.NewVector2(speed, 0);
  case movingUpLeft:
    ySpriteOffset = 3; 
    vel = rl.NewVector2(speed, speed);
  case movingUp:
    ySpriteOffset = 4;
    vel = rl.NewVector2(0, speed);
  case movingUpRight:
    ySpriteOffset = 5;
    vel = rl.NewVector2(-speed, speed);
  case movingRight:
    ySpriteOffset = 6;
    vel = rl.NewVector2(-speed, 0);
  case movingDownRight:
    ySpriteOffset = 7;
    vel = rl.NewVector2(-speed, -speed);
}

  player.src.Y = player.src.Height * ySpriteOffset;
  player.pos.X += vel.X; 
  player.pos.Y += vel.Y;

  player.src.X = player.src.Width * float32(p.spriteFrame); 

  if frameCount % 6 == 0 { 
    p.spriteFrame++; 

    // TODO, max sprite frame count should be dynamic to be asset agnositc
    if p.spriteFrame > 3 {
      player.movementState = movementStateWalking;
    }
  }
}

func (p *Player) move() {
  // TODO, vector normalization for diagonal movement
  var ySpriteOffset float32;
  var vel rl.Vector2;

  switch player.dir {
  case movingIdle:
    vel = rl.NewVector2(0, 0); 
  case movingDown: 
    ySpriteOffset = 0;
    vel = rl.NewVector2(0, -player.speed);
  case movingDownLeft: 
    ySpriteOffset = 1;
    vel = rl.NewVector2(player.speed, -player.speed);
  case movingLeft:
    ySpriteOffset = 2;
    vel = rl.NewVector2(player.speed, 0);
  case movingUpLeft:
    ySpriteOffset = 3; 
    vel = rl.NewVector2(player.speed, player.speed);
  case movingUp:
    ySpriteOffset = 4;
    vel = rl.NewVector2(0, player.speed);
  case movingUpRight:
    ySpriteOffset = 5;
    vel = rl.NewVector2(-player.speed, player.speed);
  case movingRight:
    ySpriteOffset = 6;
    vel = rl.NewVector2(-player.speed, 0);
  case movingDownRight:
    ySpriteOffset = 7;
    vel = rl.NewVector2(-player.speed, -player.speed);
}

  player.src.Y = player.src.Height * ySpriteOffset;
  player.pos.X += vel.X; 
  player.pos.Y += vel.Y;

  player.src.X = player.src.Width * float32(p.spriteFrame); 

  if frameCount % 8 == 0 { 
    p.spriteFrame++; 

    if p.spriteFrame > 3 {
      p.spriteFrame = 0; 
    }
  }

  player.setDirection(movingIdle);
}

func (p *Player) slash() {
  player.movementState = movementStateSlashing;
  // TODO, vector normalization for diagonal movement

  var ySpriteOffset float32;

  switch player.dir {
  case movingDownLeft: 
    ySpriteOffset = 0;
  case movingUpLeft, movingLeft:
    ySpriteOffset = 1; 
  case movingUpRight, movingUp:
    ySpriteOffset = 2;
  case movingDownRight, movingIdle, movingDown, movingRight:
    ySpriteOffset = 3;
  }

  player.src.Y = player.src.Height * ySpriteOffset;
  player.src.X = player.src.Width * float32(p.spriteFrame); 

  if frameCount % 6 == 0 { 
    p.spriteFrame++; 

    // TODO, max sprite frame count should be dynamic to be asset agnositc
    if p.spriteFrame > 3 {
      player.movementState = movementStateWalking;
    }
  }
}

func (p *Player) update() {
  switch player.movementState {
  case movementStateWalking: player.move(); 
  case movementStateRolling: player.roll(); 
  case movementStateSlashing: player.slash();
  case movementStateIdle: panic("unimplemented: idle state");
  }

  player.handleCollision();
}

func (p *Player) render() {
  var sprite rl.Texture2D;
  switch player.movementState {
  case movementStateIdle: panic("unimplemented: idle state");
  case movementStateWalking: sprite = player.movingSprite;
  case movementStateRolling: sprite = player.rollingSprite; 
  case movementStateSlashing: sprite = player.slashingSprite; 
  }

  rl.DrawTexturePro(sprite, player.src, player.dest, player.pos, 0, rl.RayWhite);
}

func (p *Player) unload() {
  rl.UnloadTexture(p.movingSprite); 
  rl.UnloadTexture(p.rollingSprite); 
  rl.UnloadTexture(p.slashingSprite); 
}

func (p *Player) resetAnimation() {
  player.spriteFrame = 0; 
}

type MovementDirection int;
const (
  movingIdle MovementDirection = 0b0000;
  movingDown MovementDirection = 0b0001;
  movingUp MovementDirection = 0b0010;
  movingLeft MovementDirection = 0b0100;
  movingRight MovementDirection = 0b1000;
  movingDownLeft MovementDirection = movingDown | movingLeft;
  movingUpLeft MovementDirection = movingUp | movingLeft;
  movingUpRight MovementDirection = movingUp | movingRight;
  movingDownRight MovementDirection = movingDown | movingRight;
)

var coin = rl.NewRectangle(200, 200, 50, 50); 
var coinCollected = false;

func (p *Player) handleCollision() {
  playerBox := rl.NewRectangle(-p.pos.X, -p.pos.Y, p.dest.Width, p.dest.Height);
  if rl.CheckCollisionRecs(playerBox, coin) {
    coinCollected = true; 
  } 
}

type Screen struct {
  Width int32;
  Height int32;
  Title string;
}

var (
  player *Player; 
  running = true; 
  screen = Screen{25 * 48, 25 * 48, "Ricky Boy"};
)

type TileUnit int; 

func loadMap(filePath string) *Map {
  jason, err := os.ReadFile(filePath);
  if err != nil {
    log.Fatalf("unable to load map %q", err); 
  }

  mab := new(Map); 
  err = json.Unmarshal(jason, mab);
  if err != nil {
    log.Fatalf("unable to decode map data %q", err); 
  }

  for i := range mab.Tilesets {
    mab.Tilesets[i].Texture = rl.LoadTexture("./assets/maps/" + mab.Tilesets[i].Name + ".png"); 
  }

  fmt.Printf("INFO: map %q loaded successfully", filePath);

  return mab;
}

func (m *Map) render() {
  tileDest := rl.NewRectangle(0, 0, float32(mab.TileWidth), float32(mab.TileHeight));

  for _, layer := range m.Layers {
    for i, v := range layer.Data {
      if v != 0 {

        var ts Tileset;
        for j := range mab.Tilesets {
          if v >= mab.Tilesets[j].FirstGID {
            if j + 1 == len(mab.Tilesets) {
              ts = mab.Tilesets[j];
              break;
            } 

            if v < mab.Tilesets[j+1].FirstGID {
              ts = mab.Tilesets[j];
              break;
            }
          }
        }

        tileSrc := rl.NewRectangle(float32(ts.Margin), float32(ts.Margin), float32(ts.TileWidth), float32(ts.TileHeight));

        tileDest.X = tileDest.Width * float32(i % mab.Width); 
        tileDest.Y = tileDest.Height * float32(i / mab.Width); 

        tileSrc.X = tileSrc.Width * float32((v-ts.FirstGID) % int(ts.Columns)) + float32(ts.Spacing);
        tileSrc.Y = tileSrc.Height * float32((v-ts.FirstGID) / int(ts.Columns)) + float32(ts.Spacing);

        log.Println(ts.Name, ts.Columns, tileSrc, tileDest)
        rl.DrawTexturePro(ts.Texture, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White);
      }
    }
  }
}

func (m *Map) unload() {
  for _, tileset := range mab.Tilesets {
    rl.UnloadTexture(tileset.Texture); 
  }
}

func start() {
  rl.InitWindow(screen.Width, screen.Height, screen.Title);
  loadPlayer();
  rl.SetExitKey(rl.KeyQ); 
  rl.SetTargetFPS(60); 
  mab = loadMap("./assets/maps/map.json");
}

func quit() {
  player.unload(); 
  mab.unload();
  rl.CloseWindow();
}

func loadPlayer() {
  movingSprite := rl.LoadTexture("assets/ricky/Character/ricky-moving.png");
  rollingSprite := rl.LoadTexture("assets/ricky/Character/ricky-rolling-no-pun-intended.png");
  slashingSprite := rl.LoadTexture("assets/ricky/Character/ricky-slashing.png");
  src := rl.NewRectangle(0, 0, 32, 32); 
  dest := rl.NewRectangle(0, 0, 128, 128); 
  pos := rl.NewVector2(0, 0); 

  player = NewPlayer(movingSprite, rollingSprite, slashingSprite, src, dest, pos, 3.5); 
}

func input() {
  if rl.IsKeyDown(rl.KeyW) {
    player.setDirection(movingUp);
  } else if rl.IsKeyDown(rl.KeyS) {
    player.setDirection(movingDown);
  }

  if rl.IsKeyDown(rl.KeyA) {
    player.setDirection(movingLeft);
  } else if rl.IsKeyDown(rl.KeyD) {
    player.setDirection(movingRight);
  }

  if rl.IsKeyPressed(rl.KeyC) && !player.isMidAnimation() {
    player.resetAnimation();
    player.roll();
  }

  if rl.IsKeyPressed(rl.KeyF) && !player.isMidAnimation() {
    player.resetAnimation();
    player.slash();
  }
}

func update() {
  running = !rl.WindowShouldClose(); 
  player.update(); 
}

func render() {
  rl.BeginDrawing();

  mab.render();
  rl.DrawText("Ricky in Town", 50, 50, 48, rl.NewColor(252, 176, 179, 255));

  if !coinCollected {
    rl.DrawRectangle(coin.ToInt32().X, coin.ToInt32().Y, coin.ToInt32().Width, coin.ToInt32().Height, rl.Yellow);
  }

  player.render(); 

  rl.ClearBackground(rl.NewColor(252, 236, 201, 255));
  rl.EndDrawing();

  frameCount++; 
} 

func main() {
  start();
  defer quit(); 

  for running {
    input(); 
    update();
    render(); 
  }
}
