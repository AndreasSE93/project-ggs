package clientHandlers;

import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.io.IOException;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.HostRoom;
import packageManaging.HostRoomEncoder;
import packageManaging.StageFlipper;

import snake.SnakeGUI;
import snake.SnakeMessageEncoder;
import snake.SnakeServerMessage;
import snake.SnakeUserMessage;

import clientNetworking.NetManager;

public class SnakeHandler implements HandlerInterface, KeyListener {
	NetManager network;
	final String userName;
	int Player = 2;
	SnakeMessageEncoder SME = new SnakeMessageEncoder();
	HostRoomEncoder hre = new HostRoomEncoder();
	SnakeGUI SG;
	boolean loop = true;

	public SnakeHandler(NetManager net, String username) {
		network = net;
		this.userName = username;
	}

	@Override
	public StageFlipper init(StageFlipper nothing) {
		SG = new SnakeGUI();
		SG.render();
		SG.gamePane.addKeyListener(this);

		while (loop) {
			try {
				String mess = receiveMessage();

				int id = retrieveId(mess);
				decodeAndRender(id, mess);

			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}

		return null;
	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		
		case 101: //Init Message
			
			HostRoom hr = hre.decode(message);
			if(Player == 0)
			this.Player = hr.Player;
			
			break;
			
		case 102:
			HostRoom hr2 = hre.decode(message);
			if(Player == 0)
			this.Player = hr2.Player;
			
			break;
			
		case 302: // Recieved updated movements from players, repaint board
			
			SnakeServerMessage SSM = SME.decode(message);
			SG.repaint(SSM.Players);
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
		if(keyEvent != ""){
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

}
