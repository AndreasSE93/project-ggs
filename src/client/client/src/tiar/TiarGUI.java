package tiar;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Image;
import java.awt.Insets;
import java.io.File;
import java.io.IOException;



import javax.imageio.ImageIO;
import javax.swing.BoxLayout;
import javax.swing.ImageIcon;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;

public class TiarGUI extends GameLogic{
	public JButton a1 = new JButton();
	public JButton a2 = new JButton();
	public JButton a3 = new JButton();
	public JButton b1 = new JButton();
	public JButton b2 = new JButton();
	public JButton b3 = new JButton();
	public JButton c1 = new JButton();
	public JButton c2 = new JButton();
	public JButton c3 = new JButton();
	public JButton[] game = new JButton[9];
	
	public  void render (){
		
		JFrame window = new JFrame();
		window.setSize(800, 600);
		window.setLayout(new BorderLayout());
		window.setLocationRelativeTo(null);
		window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		window.setBackground(Color.DARK_GRAY);
		
		JPanel gamePane = new JPanel();
		gamePane.setPreferredSize(new Dimension(500, 600));
		gamePane.setBackground(Color.BLACK);
		
		JPanel topRow = new JPanel();
		JPanel middleRow = new JPanel();
		JPanel bottomRow = new JPanel();
		
		
		gamePane.setLayout(new BoxLayout(gamePane, BoxLayout.Y_AXIS));
		topRow.setLayout(new BoxLayout(topRow, BoxLayout.X_AXIS));
		middleRow.setLayout(new BoxLayout(middleRow, BoxLayout.X_AXIS));
		bottomRow.setLayout(new BoxLayout(bottomRow, BoxLayout.X_AXIS));
		
		
		 JPanel aa1 = new JPanel();
		 JPanel aa2 = new JPanel();
		 JPanel aa3 = new JPanel();
		 JPanel bb1 = new JPanel();
		 JPanel bb2 = new JPanel();
		 JPanel bb3 = new JPanel();
		 JPanel cc1 = new JPanel();
		 JPanel cc2 = new JPanel();
		 JPanel cc3 = new JPanel();
		 
		 
		 a1.setName("a1");
		 a2.setName("a2");
		 a3.setName("a3");
		 a1.setName("b1");
		 b2.setName("b2");
		 b3.setName("b3");
		 c1.setName("c1");
		 c2.setName("c2");
		 c3.setName("c3");
		 
		 
		 
		 
		 aa1.add(a1);
		 aa2.add(a2);
		 aa3.add(a3);
		 bb1.add(b1);
		 bb2.add(b2);
		 bb3.add(b3);
		 cc1.add(c1);
		 cc2.add(c2);
		 cc3.add(c3);
		
		topRow.add(aa1);
		topRow.add(aa2);
		topRow.add(aa3);
		
		middleRow.add(bb1);
		middleRow.add(bb2);
		middleRow.add(bb3);
		
		bottomRow.add(cc1);
		bottomRow.add(cc2);
		bottomRow.add(cc3);
		
		topRow.setPreferredSize(new Dimension (300, 100));
		bottomRow.setPreferredSize(new Dimension (300, 100));
		middleRow.setPreferredSize(new Dimension (300, 100));
		
		
		
		a1.setPreferredSize(new Dimension(100,100));
		a2.setPreferredSize(new Dimension(100,100));
		a3.setPreferredSize(new Dimension(100,100));
		b1.setPreferredSize(new Dimension(100,100));
		b2.setPreferredSize(new Dimension(100,100));
		b3.setPreferredSize(new Dimension(100,100));
		c1.setPreferredSize(new Dimension(100,100));
		c2.setPreferredSize(new Dimension(100,100));
		c3.setPreferredSize(new Dimension(100,100));
		
		game[0] = a1;
		game[1] = a2;
		game[2] = a3;
		game[3] = b1;
		game[4] = b2;
		game[5] = b3;
		game[6] = c1;
		game[7] = c2;
		game[8] = c3;
		
		this.c1.setBackground(Color.BLUE);
		this.c2.setBackground(Color.RED);
		this.c3.setBackground(Color.GREEN);
		
		this.a1.setBackground(Color.BLUE);
		this.a2.setBackground(Color.RED);
		this.a3.setBackground(Color.GREEN);
		
		
		this.b3.setBackground(Color.BLUE);
		this.b1.setBackground(Color.RED);
		this.b2.setBackground(Color.GREEN);
		
		gamePane.add(topRow);
		gamePane.add(middleRow);
		gamePane.add(bottomRow);
		
		
		
		
		
		window.add(gamePane, BorderLayout.WEST);
		window.setVisible(true);
		
		
		
		
	}
	
	
	
	public void doMove (String position, int player) {
		
		
		int pos = getInt(position);
		if(super.validMove(pos, player)){ 
		JButton update = this.game[pos];
		makeJButton(update, player);

		}

	}
	
	private int getInt(String Pos){
		switch (Pos){
		case "a1":
			return 0;
		case "a2":
			return 1;
		case "a3":
			return 2;
		case "b1":
			return 3;
		case "b2":
			return 4;
		case "b3":
			return 5;
		case "c1":
			return 6;
		case "c2":
			return 7;
		case "c3":
			return 8;
		default:
			return 10000;
		}
		
		
	}

	private void makeJButton (JButton jb,  int player){
		File imageCheck;
		if (player == 1){ imageCheck = new File("resources/cross.png");}
		else { imageCheck = new File("resources/circle.png");
		}
		try {
			Image img = ImageIO.read(imageCheck);
			jb.setIcon(new ImageIcon(img));
		} catch (IOException ex) {
		}
		jb.setMargin(new Insets(0, 0, 0, 0));
		jb.setBorder(null);
	
		
	}
}
	
	
