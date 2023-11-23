package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var frameCount int; 

type Player struct {
  sprite rl.Texture2D;
  spriteFrame int; 
  src rl.Rectangle;
  dest rl.Rectangle; 
  pos rl.Vector2;
  dir MovingDirection;
  speed float32; 
}

func NewPlayer(sprite rl.Texture2D, src rl.Rectangle, dest rl.Rectangle, pos rl.Vector2, speed float32) *Player {
  return &Player{
    sprite: sprite, 
    src: src, 
    dest: dest, 
    pos: pos,
    speed: speed, 
  }
}

func (p *Player) setDirection(dir MovingDirection) {
  p.dir = p.dir | dir;
}

func (p *Player) move() {
  // TODO, vector normalization for diagonal movement

  var ySpriteOffset float32;
  var vel rl.Vector2;

  switch player.dir {
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
}

func (p *Player) render() {
  rl.DrawTexturePro(player.sprite, player.src, player.dest, player.pos, 0, rl.RayWhite);
  player.dir = 0;
}

type MovingDirection int;
const (
  movingIdle MovingDirection = 0b00000000;
  movingDown MovingDirection = 0b00000001;
  movingUp MovingDirection = 0b00000010;
  movingLeft MovingDirection = 0b00000100;
  movingRight MovingDirection = 0b00001000;
  movingDownLeft MovingDirection = movingDown | movingLeft;
  movingUpLeft MovingDirection = movingUp | movingLeft;
  movingUpRight MovingDirection = movingUp | movingRight;
  movingDownRight MovingDirection = movingDown | movingRight;
)

type Screen struct {
  Width int32;
  Height int32;
  Title string;
}

var (
  player *Player; 
  running = true; 
  screen = Screen{640, 480, "Ricky Boy"};
)

func start() {
  rl.InitWindow(screen.Width, screen.Height, screen.Title);
  loadPlayer();
  rl.SetExitKey(rl.KeyQ); 
  rl.SetTargetFPS(60); 
}

func quit() {
  unloadTextures(); 
  rl.CloseWindow();
}

func loadPlayer() {
  sprite := rl.LoadTexture("assets/ricky/Character/ricky-moving.png");
  src := rl.NewRectangle(0, 0, 32, 32); 
  dest := rl.NewRectangle(0, 0, 64, 64); 
  pos := rl.NewVector2(0, 0); 
  player = NewPlayer(sprite, src, dest, pos, 1.4); 
}

func unloadTextures() {
  rl.UnloadTexture(player.sprite); 
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
}

func update() {
  running = !rl.WindowShouldClose(); 
  player.move();
}

func render() {
  rl.BeginDrawing();
  rl.DrawText("Ricky in Town", 50, 50, 24, rl.RayWhite);

  player.render(); 

  rl.ClearBackground(rl.NewColor(30, 39, 73, 255));
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
