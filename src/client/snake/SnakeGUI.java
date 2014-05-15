package snake;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.Image;
import java.awt.image.BufferedImage;

import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextArea;

public class SnakeGUI {

	JFrame window = new JFrame();
	ImageIcon ii;
	BufferedImage bi;
	public JLabel gamePane;
	Image dimg;
	final private int WIDTH = 1035 - 200;
	final private int HEIGTH = 790;

	public int[] ColorArray = { Color.red.getRGB(), Color.blue.getRGB(),
			Color.green.getRGB(), Color.yellow.getRGB() };

	public void render() {
		window.setSize(1035, 790);
		window.setLayout(new BorderLayout());

		window.setLocationRelativeTo(null);
		window.setResizable(false);
		window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		window.setBackground(Color.DARK_GRAY);

		JPanel achtungPanel = new JPanel(new BorderLayout());
		achtungPanel.setPreferredSize(new Dimension(1035, 790));

		gamePane = new JLabel();
		gamePane.setFocusable(true);
		// gamePane.addKeyListener(this);

		achtungPanel.add(createTextPanel(), BorderLayout.EAST);
		achtungPanel.add(gamePane);

		bi = new BufferedImage(WIDTH, HEIGTH, BufferedImage.TYPE_INT_RGB);

		renderNewGame();

		window.add(achtungPanel, BorderLayout.NORTH);
		window.setVisible(true);
		window.validate();


		// Temp play

		/*	SnakePlayer[] players = new SnakePlayer[4];
		int posy1= 100;
		int posy2 = 200;
		int posy3 = 300;
		int posy4 = 400;
		for(int i= 0; i<(1035-200); i++){
			SnakePlayer p1 = new SnakePlayer(1,i,posy1,true);
			SnakePlayer p2 = new SnakePlayer(2,i,posy2,true);
			SnakePlayer p3 = new SnakePlayer(3,i,posy3,true);
			SnakePlayer p4 = new SnakePlayer(4,i,posy4,true);
			players[0] = p1;
			players[1] = p2;
			players[2] = p3;
			players[3] = p4;
			repaint(players);
			try {
				Thread.sleep(15);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}

		}*/


	}

	public JPanel createTextPanel() {

		JPanel gameInfo = new JPanel();
		gameInfo.setPreferredSize(new Dimension(200, 790));
		gameInfo.setBackground(Color.gray.darker());

		JTextArea gameName = new JTextArea("Achtung Die Kurve");
		Font font = new Font("Helvetica", Font.BOLD, 12);
		gameName.setFont(font);
		gameName.setEditable(false);
		gameName.setBackground(Color.gray.darker());
		gameName.setForeground(Color.WHITE);
		gameInfo.add(gameName, BorderLayout.NORTH);

		return gameInfo;

	}

	public void repaint(SnakePlayer[] players) {
		for (int i = 0; i < players.length; i++) {
			//if(i == 0){ uncomment to only track player 1
			System.out.println(i + "PosX:" + players[i].getPosX() + "PosY" + players[i].getPosY());

			if (players[i].isAlive()){


				bi.setRGB((int) players[i].getPosX(), (int) players[i].getPosY(),
						ColorArray[players[i].PlayerID - 1]);
				ii = new ImageIcon();
				ii.setImage(bi);

				gamePane.setIcon(ii);
				gamePane.validate();
				//	}
			}
		}
	}

	public void renderNewGame() {

		/*
		 * for(int index= 0; index<PlayerArray.length; index++){
		 * PlayerArray[index].revivePlayer(); }
		 */

		for (int k = 0; k < WIDTH; k++) {
			for (int i = 1; i < HEIGTH; i++) {
				bi.setRGB(k, i, Color.BLACK.getRGB());

			}
		}
		ii = new ImageIcon();
		ii.setImage(bi);

		gamePane.setIcon(ii);
		gamePane.validate();

	}


}
