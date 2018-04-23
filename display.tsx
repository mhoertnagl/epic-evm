import * as React from 'react';

export class Display extends React.Component<undefined, undefined> {

  componentDidMount() {
    const canvas = this.refs.canvas as HTMLCanvasElement
    const ctx = canvas.getContext('2d')

    if (ctx) {
      ctx.fillStyle = 'black'
      ctx.fillRect(0, 0, 640, 480)      
    }
  }

  render() {
    return (
      <div>
        <canvas 
          ref="canvas" 
          width={640} 
          height={480}>
        </canvas>
      </div>
    );
  }
}
