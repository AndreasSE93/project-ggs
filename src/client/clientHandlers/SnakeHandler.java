package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.io.IOException;

import javax.swing.JOptionPane;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.HostRoomEncoder;
import packageManaging.KickMessage;
import packageManaging.KickMessageEncoder;
import packageManaging.StageFlipper;
import packageManaging.TiarStartableMessage;
import packageManaging.TiarStartableMessageEncoder;
import packageManaging.TiarStartedMessage;
import packageManaging.TiarStartedMessageEncoder;

import snake.SnakeGUI;
import snake.SnakeMessageEncoder;
import snake.SnakeServerMessage;
import snake.SnakeUserMessage;

import clientNetworking.NetManager;

public class SnakeHandler implements HandlerInterface, KeyListener, ActionListener {
	NetManager network;
	final String userName;
	int Player = 2;
	SnakeMessageEncoder SME = new SnakeMessageEncoder();
	HostRoomEncoder hre = new HostRoomEncoder();
	KickMessageEncoder kEnc = new KickMessageEncoder();
	TiarStartableMessageEncoder tsme = new TiarStartableMessageEncoder();
	TiarStartedMessageEncoder tiStarter = new TiarStartedMessageEncoder();
	StageFlipper saveMsg;
	SnakeGUI SG;
	boolean loop = true;
	private boolean startable, started;

	public SnakeHandler(NetManager net, String username) {
		network = net;
		this.userName = username;
	}

	@Override
	public StageFlipper init(StageFlipper nothing) {
		SG = new SnakeGUI();
		SG.render();
		SG.achtungPanel.addKeyListener(this);
		SG.startGame.addActionListener(this);
		SG.startGame.addKeyListener(this);
		SG.leaveGame.addActionListener(this);
		
		while (loop) {
			try {
				String mess = receiveMessage();
				int id = retrieveId(mess);
				decodeAndRender(id, mess);

			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}

		SG.window.setVisible(false);
		return saveMsg;
	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		
		case 202:
			TiarStartableMessage tsm = tsme.decode(message);
			this.startable = tsm.isStartable;
			break;
			
		case 203:
			TiarStartedMessage startedGame = tiStarter.decode(message);
			this.started = startedGame.started;
			this.Player = startedGame.playerID;
			this.SG.achtungPanel.requestFocus();
			break;
			
		case 404:
			System.out.println("404 kcik");
			this.saveMsg = new StageFlipper(kEnc.decode(message));
			this.loop = false;
			break;
			
		case 302: // Recieved updated movements from players, repaint board
			SnakeServerMessage SSM = SME.decode(message);
			this.SG.achtungPanel.requestFocus();
			if(!SSM.clearBoard)
			SG.repaint(SSM.Players);
				
			else{
			SG.renderNewGame();
			System.out.println(SSM.hasWon);
			if (SSM.hasWon){
				JOptionPane.showMessageDialog(null, SSM.winnerName + " has won!", "WinnerWinnerChickenDinner", JOptionPane.INFORMATION_MESSAGE);
				
			}
				}
			if (SG.nameSet == 0)
			SG.setNames(SSM);
			
			break;
			
		default:
			System.out.println("Should not come here! " + id); // might be 103 refresh msg
		}
	}

	
	
	@Override
	public void sendMessage(String message) {
		
		network.send(message);

	}

	@Override
	public String receiveMessage() throws IOException {
		return network.receiveMessage();
	}

	@Override
	public int retrieveId(String mess) {
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {

		}
		return 0;

	}

	@Override
	public void keyPressed(KeyEvent e) {
		
		int key = e.getKeyCode();
		String keyEvent = "";
		if (key == KeyEvent.VK_LEFT) {
			keyEvent = "VK_LEFT";

		}
		if (key == KeyEvent.VK_RIGHT) {
			keyEvent = "VK_RIGHT";
		}
		if(keyEvent != "" && started){
		SnakeUserMessage SUM = new SnakeUserMessage(Player, keyEvent);

		try {
			sendMessage(SME.encodeSnakeUserMessage(SUM));
		} catch (JSONException e1) {
			// TODO Auto-generated catch block
			e1.printStackTrace();
		}
		}
	}

	@Override
	public void keyReleased(KeyEvent e) {
		// TODO Auto-generated method stub

	}

	@Override
	public void keyTyped(KeyEvent e) {
		// TODO Auto-generated method stub

	}

	@Override
	public void actionPerformed(ActionEvent arg0) {
		String command = arg0.getActionCommand();
		String JSONtext = "";
		switch (command) {
		case "startButton":
		if(startable){
			try {
				JSONtext = tiStarter.encode(new TiarStartedMessage());

				sendMessage(JSONtext);
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			
			
		}
		
		case "kick":
			try {
				JSONtext = kEnc.encode(new KickMessage());
				//aJOsSystem.out.println(JSONtext);
				sendMessage(JSONtext);
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
		
			
		
		
			
		default:
			break;
		
		
	 }

	}
}
