package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;
import java.io.IOException;

import javax.swing.JButton;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.TiarUserMessage;
import packageManaging.TiarUserMessageEncoder;

import tiar.TiarGUI;

import clientNetworking.NetManager;



/* Handler and initializer for a Tic Tac Toe game */

public class TiarHandler implements HandlerInterface, ActionListener, MouseListener{

	NetManager network;
	TiarGUI tg;
	
	public TiarHandler(NetManager net){
		this.network = net;
		
	}

	public void init(){
		 tg = new TiarGUI();
		 tg.render();
		 System.out.println("1");
		 for (int i = 0 ; i < tg.game.length; i++){
		 tg.game[i].addMouseListener(this);
		 
		 }
		 System.out.println("2");
		/* Create new Tiar lobby, probably with a graphical three in a row and a chat. 
		 * Add actionLsiteners to resp.. */
		
		tg.window.setVisible(true);
		
		
		while (true) {
			try {
				String mess = recieveMessage();
			
				int id = retrieveId(mess);
				decodeAndRender(id, mess);
				
			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
	}
	
	public void decodeAndRender(int id, String message) throws JSONException {
		switch(id){
		case 100: //Chat message
				break;
				
		case 200: //Move message
				TiarUserMessageEncoder s = new TiarUserMessageEncoder();
				TiarUserMessage mess = s.decode(message);
				tg.doMove(mess.Move , tg.getTurn());
				break;
				
		default: //Should not come here
				break;
		
		
		}
		
		
		
	}
	
	public void sendMessage(String message) {
		network.send(message);
		
	}

	
	public String recieveMessage() throws IOException {
		return network.receiveMessage();
	}
	
	
	public int retrieveId(String mess){
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {
			
		}
		return 0;
		
	}

	
	public void actionPerformed(ActionEvent arg0) {
		// TODO Auto-generated method stub
		
	}

	@Override
	public void mouseClicked(MouseEvent arg0) {
		JButton l =(JButton) arg0.getSource();
		String Name = l.getName();
		tg.doMove(Name, 1);
	}

	@Override
	public void mouseEntered(MouseEvent arg0) {
		// TODO Auto-generated method stub
		
	}

	@Override
	public void mouseExited(MouseEvent arg0) {
		// TODO Auto-generated method stub
		
	}

	@Override
	public void mousePressed(MouseEvent arg0) {
		// TODO Auto-generated method stub
		
	}

	@Override
	public void mouseReleased(MouseEvent arg0) {
		// TODO Auto-generated method stub
		
	}

}
